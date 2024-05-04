package server

import (
	"context"
	"fmt"
	"gin-api-template/internal/env"
	"gin-api-template/internal/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Run(env *env.Values) {
	logger, err := newLogger(env)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()

	logger.Info("Server start", zap.Any("env", env))

	db := dbConnect(env, logger)
	router := route.SetupRouter(env, db, logger)

	srv := &http.Server{
		Addr:    ":" + env.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failure start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), env.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown:", zap.Error(err))
	}
	logger.Info("Success server shutdown")

	<-ctx.Done()
	logger.Info(fmt.Sprintf("Timeout of %.0f seconds", env.ShutdownTimeout.Seconds()))

	logger.Info("Server exiting")
}

func newLogger(env *env.Values) (*zap.Logger, error) {
	logLevel := zap.InfoLevel
	if env.Debug {
		logLevel = zap.DebugLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // これを入れないとタイムスタンプになる

	return zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      env.Debug,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
}

func dbConnect(env *env.Values, logger *zap.Logger) *gorm.DB {
	username := env.DatabaseUsername
	password := env.DatabasePassword
	host := env.DatabaseHost
	database := env.Database
	port := env.DatabasePort

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database +
		"?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Fatal("Failure DB Connection:", zap.Error(err))
	}
	return db
}
