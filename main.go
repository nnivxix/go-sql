package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := GetConnection()
	defer db.Close()

	// createUser(db)
	getQueryContext(db)
	getQueryRowContext(db)
	// statementQuery(db)

}

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

func getQueryContext(db *sql.DB) {
	ctx := context.Background()

	sqlQuery := "SELECT id, name, email, created_at  FROM users ORDER BY `id` DESC LIMIT 5"

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
		var createdAt time.Time

		if err := rs.Scan(&id, &name, &email, &createdAt); err != nil {
			panic(err)
		}
		fmt.Println(id, name, "|email: ", email, "|created: ", createdAt)
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
	fmt.Printf("Show 5 users from total %d users\n", total)
}

func createUser(db *sql.DB) {
	name := "hanasa"
	email := "hanasa@hanasa.com"
	password, error := bcrypt.GenerateFromPassword([]byte("password"), 12)

	if error != nil {
		panic(error)
	}

	ctx := context.Background()
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	result, err := db.ExecContext(ctx, query, name, email, password)

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

func statementQuery(db *sql.DB) {
	ctx := context.Background()

	lastUserIdQuery := "SELECT id FROM users ORDER BY id DESC LIMIT 1"
	row := db.QueryRowContext(ctx, lastUserIdQuery)

	var lastUserId int64
	if err := row.Scan(&lastUserId); err != nil {
		panic(err)
	}

	fmt.Println("----")
	fmt.Printf("Last user id: %d\n", lastUserId)

	insertScript := "INSERT INTO posts(user_id, content) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, insertScript)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("Content %d", i+1)

		rs, err := statement.ExecContext(ctx, lastUserId, content)
		if err != nil {
			panic(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Content id => %d\n", id)
	}
}
