package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "hanasa:@tcp(localhost:3306)/cms-tz-com")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(2 * time.Minute)
	fmt.Println("Connected")
}
