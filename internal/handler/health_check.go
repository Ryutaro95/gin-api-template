package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck interface {
	HealthCheck(c *gin.Context)
}

type healthCheck struct{}

func NewHealthCheckHandler() HealthCheck {
	return &healthCheck{}
}

func (h *healthCheck) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
