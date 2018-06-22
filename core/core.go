package core

import "github.com/jmoiron/sqlx"

type Model struct {
	DB *sqlx.DB
}
