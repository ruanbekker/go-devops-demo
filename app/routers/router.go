package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-devops-demo/controllers"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.GET("/users", userController.GetUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	return r
}
