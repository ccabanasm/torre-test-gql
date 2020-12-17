package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "postgres://torre:torre1234@localhost:5432/torredb?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
