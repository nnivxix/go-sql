package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnection() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, name)
	db, err := sql.Open("mysql", sqlInfo)

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(2 * time.Minute)

	return db
}
