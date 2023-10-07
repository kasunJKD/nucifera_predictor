package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	DB *sql.DB
}

type DBConfigFlask struct {
	DB *sql.DB
}

//Connect database
func Connect(host, port, user, password, dbname string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbConn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error Connecting to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error Pinging to database", err)
	}

	log.Println("Successfully Connected to database")

	return dbConn
}

//Initialize database ::TODO
//db.InitDB(dbConn)
//db function definitions::TODO
