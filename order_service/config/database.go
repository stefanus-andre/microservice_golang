package config

import (
	"database/sql"
	"log"
)

func ConnectDatabase() (*sql.DB, error) {
	dsn := "root@tcp(127.0.0.1)/userservice"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
