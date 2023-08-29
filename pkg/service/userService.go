package service

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user backendTraineeAssignment2023.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id int) (string, error) {
	return s.repo.DeleteUser(id)
}

func (s *UserService) UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int) (bool, string, error) {
	return s.repo.UpdateUserById(slugsToJoin, slugsToLeave, id)
}
