package conn

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	Driver   string
	Database string
}

func InitDB(cfg DBConfig) *sqlx.DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.UserName, cfg.Password, cfg.Host, cfg.Database)
	db, err := sqlx.Connect(cfg.Driver, connString)
	if err != nil {
		log.Fatalln("Error to connect database")
		return nil
	}
	log.Println("Database successfully connected")
	return db
}
