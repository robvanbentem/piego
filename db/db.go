package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

var db *sqlx.DB

func InitDB(host string, port int, user string, pass string,scheme string) {
	var err error

	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, scheme)
	db, err = sqlx.Connect("mysql", str)

	if err != nil {
		panic("cannot connect to database")
	}
}

func CloseDB() {
	db.Close()
}
