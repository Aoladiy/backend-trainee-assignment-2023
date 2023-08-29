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
	if exists, err := r.segmentExists(segment.Slug); err != nil {
		return "", err
	} else if exists {
		return fmt.Sprintf("there's already segment with this slug='%v'", segment.Slug), nil
	}

	query := fmt.Sprintf("INSERT INTO %v (slug) VALUES (?)", segmentsTable)
	_, err := r.db.Exec(query, segment.Slug)
	if err != nil {
		return "", err
	}

	return segment.Slug, nil
}

func (r *SegmentRepository) DeleteSegment(slug string) (bool, string, error) {
	if exists, err := r.segmentExists(slug); err != nil {
		return false, "", err
	} else if !exists {
		return false, fmt.Sprintf("there's no segment with this slug='%v'", slug), nil
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE slug=?", segmentsTable)
	_, err := r.db.Exec(query, slug)
	if err != nil {
		return false, "", err
	}

	return true, slug, nil
}
func (r *SegmentRepository) segmentExists(slug string) (bool, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE slug=?", segmentsTable)
	var count int
	err := r.db.Get(&count, checkQuery, slug)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
