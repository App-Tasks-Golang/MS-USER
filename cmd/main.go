package main

import (
	"User-Service-Go/internal/config"
	"User-Service-Go/pkg/adapters"
	"User-Service-Go/pkg/service"
	"User-Service-Go/transport"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	DB, err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("error al obtener la conexión a la base de datos: %v", err)
	}

	authClient := adapters.NewAuthClient("http://auth-service:8084")

	userRepo := adapters.NewUserRepository(DB)
	userService := services.NewUserService(userRepo, authClient)

	transport.SetupRoutes(r, userService)


	r.Run(":8082")
}