package repository

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user backendTraineeAssignment2023.User) (int, error)
	DeleteUser(id int) (string, error)
	UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int) (bool, string, error)
}

type Segment interface {
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
