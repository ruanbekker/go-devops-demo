package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruanbekker/go-devops-demo/app/controllers"
	"github.com/ruanbekker/go-devops-demo/app/middleware"
)

func SetupRouter(userController *controllers.UserController, healthController *controllers.HealthController) *gin.Engine {
	r := gin.Default()

	// Prometheus Middleware
	r.Use(middleware.PrometheusMiddleware())

	// API Routes
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Healthcheck Routes
	r.GET("/-/health/ready", healthController.CheckHealth)

	// Prometheus metrics endpoint
	r.GET("/metrics", middleware.PrometheusHandler())

	return r
}
