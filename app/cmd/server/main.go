package main

import (
	"os"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-devops-demo/config"
	"github.com/go-devops-demo/controllers"
	"github.com/go-devops-demo/repositories"
	"github.com/go-devops-demo/routers"
	"github.com/go-devops-demo/services"
)

func main() {
	config.ConnectDB()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	r := routers.SetupRouter(userController)

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
