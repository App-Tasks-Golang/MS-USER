package services

import (
	"User-Service-Go/pkg/adapters"
	"User-Service-Go/pkg/domain"
	"fmt"
)

type UserService struct {
	UserRepo   adapters.UserRepository
	AuthClient *adapters.AuthClient
}

func NewUserService(userRepo adapters.UserRepository, authClient *adapters.AuthClient) *UserService {
	return &UserService{UserRepo: userRepo, AuthClient: authClient}
}

func (s *UserService) CreateUser(user *domain.UserReference) error {
	userSend := domain.UserRequest{
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}

	response, resp, err := s.AuthClient.RegisterUser(userSend)
	if err != nil {
		return fmt.Errorf("error calling auth service: %w", err)
	}

	if !response {
		return fmt.Errorf("auth service error: %s", resp)
	}

	userCreate := &domain.User{
		ID:       user.ID,
		Username: user.Username,
	}

	err = s.UserRepo.CreateUser(userCreate)
	if err != nil {
		return fmt.Errorf("error saving user to database: %w", err)
	}

	return nil
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.UserRepo.GetAllUsers()
}



func (s *UserService) LoginUser(username, password string) (string, error) {
	dataUser := domain.UserRequest{
		Username: username,
		Password: password,
	}

	token, err := s.AuthClient.LoginUser(dataUser)
	if err != nil {
		return "", err
	}

	return token, nil
}



func (s *UserService) LogoutUser(token string) error {
	err := s.AuthClient.LogoutUser(token)
	if err != nil {
		return fmt.Errorf("error logging out from auth service: %w", err)
	}
	return nil
}