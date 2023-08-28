package service

import "github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"

type User interface {
}

type Segment interface {
}
type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
