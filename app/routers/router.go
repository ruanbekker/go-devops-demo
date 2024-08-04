package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruanbekker/go-devops-demo/controllers"
)

func SetupRouter(userController *controllers.UserController, healthController *controllers.HealthController) *gin.Engine {
	r := gin.Default()

	// API Routes
	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	// Healthcheck Routes
	r.GET("/-/health/ready", healthController.CheckHealth)

	return r
}
