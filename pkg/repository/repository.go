package repository

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetAllUsers() ([]backendTraineeAssignment2023.User, error)
	GetUserById(id int) (bool, backendTraineeAssignment2023.User, error)
	CreateUser(user backendTraineeAssignment2023.User) (int, error)
	DeleteUser(id int) (string, error)
	UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int) (bool, string, error)
	GetUserSegments(id int) ([]string, error)
}

type Segment interface {
	GetAllSegments() ([]backendTraineeAssignment2023.Segment, error)
	GetSegmentBySlug(slug string) (bool, backendTraineeAssignment2023.Segment, error)
	CreateSegment(segment backendTraineeAssignment2023.Segment) (string, error)
	DeleteSegment(slug string) (bool, string, error)
}
type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Segment: NewSegmentRepository(db),
	}
}
