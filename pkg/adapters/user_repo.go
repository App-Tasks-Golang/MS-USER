package adapters

import (
	"User-Service-Go/pkg/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}