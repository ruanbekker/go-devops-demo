package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ruanbekker/go-devops-demo/controllers"
	"github.com/ruanbekker/go-devops-demo/models"
	"github.com/ruanbekker/go-devops-demo/repositories"
	"github.com/ruanbekker/go-devops-demo/services"
	"github.com/ruanbekker/go-devops-demo/config"
)

func setupRouter() *gin.Engine {
	config.ConnectDB()
	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	return routers.SetupRouter(userController)
}

func TestGetUsers(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]")
}
