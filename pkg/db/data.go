package db

import (
	"database/sql"
	"log"
)

var (
	data *Data
	once bool = false
)

type Data struct {
	DB *sql.DB
}

func initDb() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	if !once {
		once = true
		initDb()
	}

	return data
}
