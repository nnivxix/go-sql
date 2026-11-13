package repository

import (
	"context"
	"go_sql/entity"
)

type UserRepositoryInf interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id int32) (entity.User, error)
}
