package main

import (
	"context"
	"database/sql"
	"fmt"
	"go_sql/connection"
	"go_sql/entity"
	"go_sql/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connection.GetConnection()
	defer db.Close()

	getAllUser(db)
	// getUser(db)
	// createUser(db)
}

func getAllUser(db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	ctx := context.Background()

	users, err := userRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%d - %s \n", user.Id, user.Email)
	}

}

func getUser(db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	ctx := context.Background()

	user, err := userRepository.FindById(ctx, 8)

	if err != nil {
		panic(err)
	}

	fmt.Printf("User found with email: %s \n", user.Email)
}

func createUser(db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	ctx := context.Background()

	user := entity.User{
		Name:     "Test",
		Email:    "test@mail.com",
		Password: "password",
	}
	user, err := userRepository.Insert(ctx, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
