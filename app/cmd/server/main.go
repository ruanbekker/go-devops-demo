package main

import (
	"os"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruanbekker/go-devops-demo/config"
	"github.com/ruanbekker/go-devops-demo/controllers"
	"github.com/ruanbekker/go-devops-demo/repositories"
	"github.com/ruanbekker/go-devops-demo/routers"
	"github.com/ruanbekker/go-devops-demo/services"
)

func main() {
	config.ConnectDB()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	healthController := controllers.NewHealthController()

	r := routers.SetupRouter(userController, healthController)

	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	if trustedProxies != "" {
		proxyList := strings.Split(trustedProxies, ",")
		err := r.SetTrustedProxies(proxyList)
		if err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	} else {
		err := r.SetTrustedProxies(nil)
		if err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	}

	r.Run(":8080")
}
