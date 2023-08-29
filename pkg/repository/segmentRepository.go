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
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE slug=?", segmentsTable)
	var count int
	err := r.db.Get(&count, checkQuery, segment.Slug)
	if err != nil {
		return "", err
	}

	if count > 0 {
		return fmt.Sprintf("segment with this slug='%v' already exists", segment.Slug), nil
	}

	query := fmt.Sprintf("INSERT INTO %v (slug) VALUES (?)", segmentsTable)
	_, err = r.db.Exec(query, segment.Slug)
	if err != nil {
		return "", err
	}

	return segment.Slug, nil
}

func (r *SegmentRepository) DeleteSegment(slug string) (bool, string, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE slug=?", segmentsTable)
	var count int
	err := r.db.Get(&count, checkQuery, slug)
	if err != nil {
		return false, "", err
	}

	if count == 0 {
		return false, fmt.Sprintf("there's no segment with this slug='%v'", slug), nil
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE slug=?", segmentsTable)
	_, err = r.db.Exec(query, slug)
	if err != nil {
		return false, "", err
	}

	return true, slug, nil
}
