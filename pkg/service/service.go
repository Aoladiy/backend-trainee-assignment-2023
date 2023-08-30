package service

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
	"time"
)

type User interface {
	GetAllUsers() ([]backendTraineeAssignment2023.User, error)
	GetUserById(id int) (bool, backendTraineeAssignment2023.User, error)
	CreateUser(user backendTraineeAssignment2023.User) (int, error)
	DeleteUser(id int) (string, error)
	UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int, ttl time.Duration) (bool, string, error)
	GetUserSegments(id int) ([]string, error)
	GetUserLog(id int, period string) (bool, []backendTraineeAssignment2023.LogEntry, error)
}

type Segment interface {
	GetAllSegments() ([]backendTraineeAssignment2023.Segment, error)
	GetSegmentBySlug(slug string) (bool, backendTraineeAssignment2023.Segment, error)
	CreateSegment(segment backendTraineeAssignment2023.Segment, autoAssignPercentage int) (string, error)
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
