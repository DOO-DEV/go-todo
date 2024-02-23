package healthservice

import (
	"database/sql"
)

type Service struct {
	healthToken string

	db *sql.DB
}

func New(db *sql.DB, token string) Service {
	return Service{
		healthToken: token,
		db:          db,
	}
}
