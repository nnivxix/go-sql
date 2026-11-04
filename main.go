package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := GetConnection()
	defer db.Close()

	getQueryContext(db)
	getQueryRowContext(db)

}

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "hanasa:@tcp(localhost:3306)/cms-tz-com")

	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(2 * time.Minute)

	return db
}

func getQueryContext(db *sql.DB) {
	ctx := context.Background()

	sqlQuery := "SELECT id, name, email FROM users LIMIT 5"
	rs, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

	fmt.Println("Showing 5 users")
	fmt.Println("----")
	defer rs.Close()

	for rs.Next() {
		var id int64
		var name string
		var email string

		if err := rs.Scan(&id, &name, &email); err != nil {
			panic(err)
		}
		fmt.Println(id, name, "---", email)
	}
	if err := rs.Err(); err != nil {
		panic(err)
	}

}

func getQueryRowContext(db *sql.DB) {
	var total int
	ctx := context.Background()
	query := "SELECT COUNT(id) FROM users"
	row := db.QueryRowContext(ctx, query)

	if row.Err() != nil {
		panic(row.Err())
	}
	row.Scan(&total)

	fmt.Println("----")
	fmt.Printf("Show 5 users from total %d users", total)
}
