package repository

import (
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/jmoiron/sqlx"
)

type SegmentRepository struct {
	db *sqlx.DB
}

func NewSegmentRepository(db *sqlx.DB) *SegmentRepository {
	return &SegmentRepository{db: db}
}

func (r *SegmentRepository) CreateSegment(segment backendTraineeAssignment2023.Segment) (string, error) {
	query := fmt.Sprintf("INSERT INTO %v (slug) VALUES (?)", segmentsTable)
	_, err := r.db.Exec(query, segment.Slug)
	if err != nil {
		return "", err
	}

	return segment.Slug, nil
}

func (r *SegmentRepository) DeleteSegment(slug string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %v WHERE slug=?", segmentsTable)
	_, err := r.db.Exec(query, slug)
	if err != nil {
		return "", err
	}

	return slug, nil
}
