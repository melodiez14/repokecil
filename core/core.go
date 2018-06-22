package core

import "github.com/jmoiron/sqlx"

// Model is struct which used to make a composition of core
// In OOP, this mechanism similar with inheritance
type Model struct {
	DB *sqlx.DB
}
