package repository

import (
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
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE id=?", usersTable)
	var count int
	err := r.db.Get(&count, checkQuery, id)
	if err != nil {
		return "", err
	}

	if count == 0 {
		return fmt.Sprintf("there's no user with this id='%v'", id), nil
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE id=?", usersTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(id), nil
}
