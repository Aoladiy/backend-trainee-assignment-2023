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

func (r *SegmentRepository) GetAllSegments() ([]backendTraineeAssignment2023.Segment, error) {
	var segments []backendTraineeAssignment2023.Segment
	query := fmt.Sprintf("SELECT * FROM %v", segmentsTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var segment backendTraineeAssignment2023.Segment
		var segmentSlug string
		var segmentID int
		err := rows.Scan(&segmentID, &segmentSlug)
		if err != nil {
			return nil, err
		}
		segment.Slug = segmentSlug
		segments = append(segments, segment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}

func (r *SegmentRepository) GetSegmentBySlug(slug string) (bool, backendTraineeAssignment2023.Segment, error) {
	if exists, err := r.segmentExists(slug); err != nil {
		return false, backendTraineeAssignment2023.Segment{}, err
	} else if !exists {
		return false, backendTraineeAssignment2023.Segment{}, nil
	}

	var segment backendTraineeAssignment2023.Segment
	query := fmt.Sprintf("SELECT * FROM %v WHERE slug=?", segmentsTable)
	err := r.db.Get(&segment, query, slug)
	if err != nil {
		return false, segment, err
	}

	return true, segment, nil
}

func (r *SegmentRepository) CreateSegment(segment backendTraineeAssignment2023.Segment, autoAssignPercentage int) (string, error) {
	if exists, err := r.segmentExists(segment.Slug); err != nil {
		return "", err
	} else if exists {
		return fmt.Sprintf("there's already a segment with this slug='%v'", segment.Slug), nil
	}

	query := fmt.Sprintf("INSERT INTO %v (slug) VALUES (?)", segmentsTable)
	_, err := r.db.Exec(query, segment.Slug)
	if err != nil {
		return "", err
	}
	if autoAssignPercentage > 0 && autoAssignPercentage <= 100 {
		segmentID, err := r.getSegmentIDBySlug(segment.Slug)
		if err != nil {
			return "", err
		}
		_, err = r.getAutoAssignedUsers(segmentID, autoAssignPercentage)
		if err != nil {
			return "", err
		}
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
func (r *SegmentRepository) getSegmentIDBySlug(slug string) (int, error) {
	var segmentID int
	var query string

	query = fmt.Sprintf("SELECT id FROM %v WHERE slug = ?", segmentsTable)

	row := r.db.QueryRow(query, slug)
	if err := row.Scan(&segmentID); err != nil {
		return 0, err
	}

	return segmentID, nil
}
func (r *SegmentRepository) getAutoAssignedUsers(segmentID int, autoAssignPercentage int) ([]int, error) {
	var userIDs []int

	totalUserCountQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v", usersTable)
	var totalUserCount int
	err := r.db.Get(&totalUserCount, totalUserCountQuery)
	if err != nil {
		return []int{}, err
	}
	numUsersToAutoAssign := int(float64(autoAssignPercentage) / 100 * float64(totalUserCount))

	query := fmt.Sprintf("SELECT id FROM %v ORDER BY RAND() LIMIT ?", usersTable)
	rows, err := r.db.Query(query, numUsersToAutoAssign)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID int
		if err = rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
		query := fmt.Sprintf("INSERT INTO %v (user_id, segment_id) VALUES (?, ?)", segmentsUsersTable)
		_, err := r.db.Exec(query, userID, segmentID)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}
