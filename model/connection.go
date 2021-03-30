package model

import (
	"fmt"
	"database/sql"
)

const (
	host			= "localhost"
	port 			= 5432
	user 			= "postgres"
	password 	= "1234"
	dbname 		= "super"
)

func Connection () *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
