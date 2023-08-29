package service

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
)

type User interface {
	CreateUser(user backendTraineeAssignment2023.User) (int, error)
	DeleteUser(id int) (string, error)
}

type Segment interface {
	CreateSegment(segment backendTraineeAssignment2023.Segment) (string, error)
	DeleteSegment(slug string) (bool, string, error)
}
type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repos.User),
		Segment: NewSegmentService(repos.Segment),
	}
}
