package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	PostgresDB struct {
		DBLive *sqlx.DB
	}
)

// var schema = `
// CREATE TABLE person (
//     first_name text,
//     last_name text,
//     email text
// );

// CREATE TABLE place (
//     country text,
//     city text NULL,
//     telcode integer
// )`

func NewPostgresDB() *PostgresDB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUsername, DBPassword, DBName)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("error cause: %+v\n", err)
		panic(err)
	}
	// db.MustExec(schema)

	err = db.Ping()
	if err != nil {
		log.Printf("error cause: %+v\n", err)
		panic(err)
	}
	fmt.Println("successfully connect to postgress")
	return &PostgresDB{
		DBLive: db,
	}
}
