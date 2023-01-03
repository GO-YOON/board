package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sqlx.DB

func New() {

	dsn := "root:1234@tcp(localhost:3306)/shop"

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatalln("CAN NOT CONNECT")
	}

	DB = db
}

