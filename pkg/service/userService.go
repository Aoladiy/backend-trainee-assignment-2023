package service

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]backendTraineeAssignment2023.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserById(id int) (bool, backendTraineeAssignment2023.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user backendTraineeAssignment2023.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id int) (string, error) {
	return s.repo.DeleteUser(id)
}

func (s *UserService) UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int, ttl time.Duration) (bool, string, error) {
	return s.repo.UpdateUserById(slugsToJoin, slugsToLeave, id, ttl)
}
func (s *UserService) GetUserSegments(id int) ([]string, error) {
	return s.repo.GetUserSegments(id)
}
func (s *UserService) GetUserLog(id int, period string) (bool, []backendTraineeAssignment2023.LogEntry, error) {
	return s.repo.GetUserLog(id, period)
}
