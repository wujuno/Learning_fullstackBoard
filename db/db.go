package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Init() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	DB = db

	fmt.Println("DB에 연결되었습니다.🔗")
	return db, nil

}