package service

import (
	"errors"
	"mongo-auth-api/internal/config"
	"mongo-auth-api/internal/models"
	"mongo-auth-api/internal/repository"
	"mongo-auth-api/pkg/utils"
)

type AuthService struct {
	repo   *repository.UserRepository
	config *config.Config
}

func NewAuthService(repo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, config: cfg}
}

func (s *AuthService) Register(username, password, data string) error {
	existingUser, _ := s.repo.FindByUsername(username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &models.User{
		Username: username,
		Password: hashedPassword,
		Data:     data,
	}

	return s.repo.CreateUser(newUser)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID.Hex(), s.config.JWTSecret)
}
