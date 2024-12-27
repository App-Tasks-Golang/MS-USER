package transport

import (
	"User-Service-Go/pkg/domain"
	"User-Service-Go/pkg/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Crear un nuevo usuario
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user *domain.UserReference
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Obtener todos los usuarios
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Obtener un usuario por su ID
func (h *UserHandler) GetOneUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Función para iniciar sesión
func (h *UserHandler) LoginUser(c *gin.Context) {
	var credentials struct {
		Username string `gorm:"unique;not null" json:"username"`
		Password string `gorm:"not null" json:"password"`
	}

	// Deserializar el JSON de la solicitud
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar credenciales con el servicio
	token, err := h.UserService.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa con el token
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"token":    token,
	})
}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	// Obtener el token desde el encabezado Authorization
	tokenString := c.GetHeader("Authorization")
	fmt.Println("Token recibido:", tokenString) // Log para depurar el token

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token encontrado"})
		return
	}

	// Validar y procesar el token
	if len(tokenString) <= len("Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido"})
		return
	}
	tokenString = tokenString[len("Bearer "):]   // Quitar el prefijo "Bearer "
	fmt.Println("Token procesado:", tokenString) // Otro log para verificar el token

	// Llamar al servicio de logout
	if err := h.UserService.LogoutUser(tokenString); err != nil {
		fmt.Println("Error en LogoutUser:", err) // Log para depurar
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo cerrar la sesión"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Sesión cerrada correctamente"})
}
