package repository

import (
	"context"
	"database/sql"
	"go-gin-repository/entity"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userMysql struct {
	DB *sql.DB
}

func NewUserMysql(db *sql.DB) UserRepository {
	return &userMysql{DB: db}
}

func (repo *userMysql) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user.ObjectID = primitive.NewObjectID()
	user.Id = user.ObjectID.Hex()
	script := "INSERT INTO users(id, name, isActive) VALUES (?, ?, ?)"
	_, err := repo.DB.ExecContext(ctx, script, user.ObjectID.Hex(), user.Name, user.IsActive)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userMysql) GetUsers(ctx context.Context) ([]entity.User, error) {
	script := "SELECT id, name, isActive FROM users"
	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.Id, &user.Name, &user.IsActive)
		users = append(users, user)
	}

	if users == nil {
		users = []entity.User{}
	}
	return users, nil
}
