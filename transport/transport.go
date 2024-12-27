package transport


import (
	"User-Service-Go/pkg/service"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(router *gin.Engine, userService *services.UserService) {
	// Crear el handler de usuario, que usa el servicio
	userHandler := NewUserHandler(userService)

	// Definir las rutas de usuarios
	users := router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.POST("/login", userHandler.LoginUser)
		users.GET("/", userHandler.GetAllUsers)
		users.GET("/:userID", userHandler.GetOneUser)
		users.POST("/logout", userHandler.LogoutUser)
	}
}
