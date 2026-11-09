package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := GetConnection()
	defer db.Close()

	// createUser(db)
	getQueryContext(db)
	getQueryRowContext(db)

}

func GetConnection() *sql.DB {

	db, err := sql.Open("mysql", "hanasa:@tcp(localhost:3306)/cms-tz-com?parseTime=true")

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

	sqlQuery := "SELECT id, name, email, email_verified_at, created_at  FROM users LIMIT 5"
	rs, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

	fmt.Println("Showing 5 users")
	fmt.Println("----")
	defer rs.Close()

	for rs.Next() {
		var id int64
		var name, email string
		var emailVerifiedAt sql.NullTime
		var createdAt time.Time

		if err := rs.Scan(&id, &name, &email, &emailVerifiedAt, &createdAt); err != nil {
			panic(err)
		}
		fmt.Println(id, name, "|email: ", email, "|verified: ", emailVerifiedAt.Time, "|created: ", createdAt)
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

func createUser(db *sql.DB) {
	password, error := bcrypt.GenerateFromPassword([]byte("password"), 12)

	if error != nil {
		panic(error)
	}

	ctx := context.Background()
	query := fmt.Sprintf("INSERT INTO users(name, email, username, password) VALUES ('hanasa2', 'hanasa2@hanasa.com', 'hanasa2', '%s')", string(password))
	result, err := db.ExecContext(ctx, query)

	if err != nil {
		fmt.Println("Duplicate Entry")
		// return
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted id=%d, rows affected=%d\n", id, affected)
}
