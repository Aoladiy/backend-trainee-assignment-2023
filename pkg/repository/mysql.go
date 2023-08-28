package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable         = "users"
	segmentsTable      = "segments"
	segmentsUsersTable = "segments_users"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewMysqlDB(cfg Config) (*sqlx.DB, error) {
	//"test:test@(localhost:3306)/test"
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%v:%v@(%v:%v)/%v",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
