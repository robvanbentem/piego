package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Connect("mysql", "piego@/piego")

	if err != nil {
		panic("cannot connect to database")
	}
}

func CloseDB() {
	db.Close()
}
