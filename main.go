package main

import (
	"go-postgres/config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	getConnection := config.NewPostgresDB()
	getConnection.DBLive.Ping()
}
