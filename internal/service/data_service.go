package service

import (
	"errors"
	"mongo-auth-api/internal/models"
	"mongo-auth-api/internal/repository"
)

type DataService struct {
	repo *repository.UserRepository
}

func NewDataService(repo *repository.UserRepository) *DataService {
	return &DataService{repo: repo}
}

func (s *DataService) GetData(userID string) (string, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return "", errors.New("user not found")
	}
	return user.Data, nil
}

func (s *DataService) GetAllUserStats() ([]models.DataMetadata, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var stats []models.DataMetadata
	for _, u := range users {
		stats = append(stats, models.DataMetadata{
			UserID:   u.Username,
			DataSize: len(u.Data),
			Preview:  u.Data,
		})
	}
	return stats, nil
}
