package config

import "os"

var (
	// DB PORT
	DBPort string = os.Getenv("DBPORT")
	// DB username
	DBUsername string = os.Getenv("DBUSERNAME")
	// DB password
	DBPassword string = os.Getenv("DBPASSWORD")
	// DB Name
	DBName string = os.Getenv("DBNAME")
	//DB Host
	DBHost string = os.Getenv("DBHOST")
)
