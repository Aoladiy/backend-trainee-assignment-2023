package service

import (
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment backendTraineeAssignment2023.Segment) (string, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(slug string) (bool, string, error) {
	return s.repo.DeleteSegment(slug)
}