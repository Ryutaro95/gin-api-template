package route

import (
	"gin-api-template/internal/env"
	"gin-api-template/internal/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(env *env.Values, db *gorm.DB, logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	// DI
	healthCheckHandler := handler.NewHealthCheckHandler()

	r.GET("/health", healthCheckHandler.HealthCheck)

	return r
}
