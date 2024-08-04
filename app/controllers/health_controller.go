package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanbekker/go-devops-demo/app/config"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) CheckHealth(ctx *gin.Context) {
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Printf("Health check failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy", "error": err.Error()})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("Health check ping failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
