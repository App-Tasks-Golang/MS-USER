package config

import (
	"User-Service-Go/pkg/domain"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"  // Importar el driver de MySQL
	"gorm.io/gorm"
)

// ConnectToDB establece la conexión con la base de datos MySQL y realiza la migración.
func ConnectToDB() (*gorm.DB, error) {
	// Cargar variables del archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando el archivo .env: %w", err)
	}

	// Formar la cadena de conexión a la base de datos para MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("USER_ROOT"),        // Usuario de la base de datos
		os.Getenv("USER_PASSWORD"),    // Contraseña de la base de datos
		os.Getenv("USER_HOST"),        // Dirección del host de la base de datos
		os.Getenv("USER_PORT"),        // Puerto de la base de datos
		os.Getenv("USER_NAME"),        // Nombre de la base de datos
	)

	// Intentar abrir la conexión a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  // Cambiar a MySQL
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	// Realizar la migración de la estructura User
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, fmt.Errorf("error al realizar la migración: %w", err)
	}

	// Devolver la conexión exitosa
	return db, nil
}
