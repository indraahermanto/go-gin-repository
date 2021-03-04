package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var MySQL *sql.DB

func ConnectMysqlDB() {
	var (
		DBHost     = os.Getenv("MYSQL_HOST")
		DBPort     = os.Getenv("MYSQL_PORT")
		DBUsername = os.Getenv("MYSQL_USERNAME")
		DBPassword = os.Getenv("MYSQL_PASSWORD")
		DBName     = os.Getenv("MYSQL_DATABASE")
	)

	// connect to the database
	fmt.Println("Connecting to Mysql")
	dbUrl := DBUsername + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"

	fmt.Println(dbUrl)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	MySQL = db
	fmt.Println("Connected to Mysql")
}
