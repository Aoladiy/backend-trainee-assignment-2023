package repository

import (
	"fmt"
	backendTraineeAssignment2023 "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user backendTraineeAssignment2023.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %v (id) VALUES (?)", usersTable)
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (r *UserRepository) DeleteUser(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %v WHERE id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
