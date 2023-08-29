package repository

import (
	"database/sql"
	"errors"
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]backendTraineeAssignment2023.User, error) {
	var users []backendTraineeAssignment2023.User
	query := fmt.Sprintf("SELECT * FROM %v", usersTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user backendTraineeAssignment2023.User
		var userId int
		err := rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
		user.Id = userId
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserById(id int) (bool, backendTraineeAssignment2023.User, error) {
	if exists, err := r.userExists(id); err != nil {
		return false, backendTraineeAssignment2023.User{}, err
	} else if !exists {
		return false, backendTraineeAssignment2023.User{}, nil
	}

	var user backendTraineeAssignment2023.User
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=?", usersTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		return false, user, err
	}

	return true, user, nil
}

func (r *UserRepository) GetUserSegments(id int) ([]string, error) {
	var segments []string

	if exists, err := r.userExists(id); err != nil {
		return []string{""}, err
	} else if !exists {
		return []string{fmt.Sprintf("there's no user with this id='%v'", id)}, nil
	}

	query := fmt.Sprintf("SELECT s.slug FROM %v s JOIN %v su ON s.id = su.segment_id WHERE su.user_id = ?", segmentsTable, segmentsUsersTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var segmentSlug string
		err := rows.Scan(&segmentSlug)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segmentSlug)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return segments, nil
}

func (r *UserRepository) CreateUser(user backendTraineeAssignment2023.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %v VALUES ()", usersTable)
	_, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}

	lastUserIdQuery := fmt.Sprintf("SELECT LAST_INSERT_ID() FROM %v LIMIT 1", usersTable)
	var id int
	err = r.db.Get(&id, lastUserIdQuery)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *UserRepository) DeleteUser(id int) (string, error) {
	if exists, err := r.userExists(id); err != nil {
		return "", err
	} else if !exists {
		return fmt.Sprintf("there's no user with this id='%v'", id), nil
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(id), nil
}

func (r *UserRepository) UpdateUserById(slugsToJoin []string, slugsToLeave []string, id int) (bool, string, error) {
	if exists, err := r.userExists(id); err != nil {
		return false, "fail", err
	} else if !exists {
		return false, fmt.Sprintf("there's no user with this id='%v'", id), nil
	}

	if len(slugsToJoin) == 0 && len(slugsToLeave) == 0 {
		return true, "success (Nothing changed, did you really wanted it?)", nil
	}
	tx, err := r.db.Begin()
	if err != nil {
		return false, "fail", err
	}

	for _, slug := range slugsToJoin {
		segmentID, err := r.getSegmentIDBySlug(tx, slug)
		if err != nil {
			tx.Rollback()
			return false, "fail", err
		}

		if segmentID == 0 {
			tx.Rollback()
			return false, fmt.Sprintf("there's no segments to join with slug='%v'", slug), nil
		}

		exists, err := r.checkSegmentUserAssociation(tx, id, segmentID)
		if err != nil {
			tx.Rollback()
			return false, "fail", err
		}

		if !exists {
			query := fmt.Sprintf("INSERT INTO %v (user_id, segment_id) VALUES (?, ?)", segmentsUsersTable)
			_, err = tx.Exec(query, id, segmentID)
			if err != nil {
				tx.Rollback()
				return false, "fail", err
			}
		}
	}

	for _, slug := range slugsToLeave {
		segmentID, err := r.getSegmentIDBySlug(tx, slug)
		if err != nil {
			tx.Rollback()
			return false, "fail", err
		}

		if segmentID == 0 {
			tx.Rollback()
			return false, fmt.Sprintf("there's no segments to leave with slug='%v'", slug), nil
		}

		query := fmt.Sprintf("DELETE FROM %v WHERE user_id = ? AND segment_id = ?", segmentsUsersTable)
		_, err = tx.Exec(query, id, segmentID)
		if err != nil {
			tx.Rollback()
			return false, "fail", err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return false, "fail", err
	}

	return true, "success", nil
}

func (r *UserRepository) getSegmentIDBySlug(tx *sql.Tx, slug string) (int, error) {
	var segmentID int
	err := tx.QueryRow("SELECT id FROM segments WHERE slug = ?", slug).Scan(&segmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return segmentID, nil
}

func (r *UserRepository) checkSegmentUserAssociation(tx *sql.Tx, userID, segmentID int) (bool, error) {
	var exists bool
	err := tx.QueryRow("SELECT EXISTS (SELECT 1 FROM segments_users WHERE user_id = ? AND segment_id = ?)", userID, segmentID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRepository) userExists(id int) (bool, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE id=?", usersTable)
	var count int
	err := r.db.Get(&count, checkQuery, id)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
