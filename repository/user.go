package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_sql/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInf {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	sql := "INSERT INTO users(name, email, password) values(?, ?, ?)"
	rs, err := repo.db.ExecContext(ctx, sql, user.Name, user.Email, user.Password)
	if err != nil {
		return user, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return user, err
	}

	fmt.Println("Insert user with id: ", id)
	user.Id = int32(id)

	return user, nil

}
func (repo *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	sql := "SELECT id, name, email FROM users"
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
func (repo *UserRepository) FindById(ctx context.Context, id int32) (entity.User, error) {
	user := entity.User{}

	sql := "SELECT id, name, email FROM users WHERE id = ? LIMIT 1"
	rows, err := repo.db.QueryContext(ctx, sql, id)

	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
		return user, nil
	}

	message := fmt.Sprintf("User with %d not found", id)

	return user, errors.New(message)
}
