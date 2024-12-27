package domain

//Estructura principal, usada solo para obtener todos los datos
type UserReference struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
}

//Estructura definitiva de User, usada para crear el usuario
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
}

//Estructura auxiliar para el request al ms auth
type UserRequest struct {
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Username string `gorm:"unique;not null" json:"username"`
}
