package repository

import (
	"context"
	"go-gin-repository/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
}
