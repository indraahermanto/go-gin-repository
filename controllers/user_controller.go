package controllers

import (
	"context"
	"fmt"
	"go-gin-repository/database"
	"go-gin-repository/entity"
	"go-gin-repository/repository"
	"go-gin-repository/shared"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func getRepo() repository.UserRepository {
	var repo repository.UserRepository
	if os.Getenv("DB_CONNECTION") == "mongodb" {
		repo = repository.NewUserMongoDB(database.MongoDB)
	} else {
		repo = repository.NewUserMysql(database.MySQL)
	}
	return repo
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// Validasi
	type reqUser struct {
		Name string `json:"name" binding:"required"`
	}
	var userInput reqUser

	if err := c.ShouldBindJSON(&userInput); err != nil {
		shared.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	// Create User
	repo := getRepo()
	user, err := repo.Create(context.Background(), entity.User{Name: userInput.Name})
	if err != nil {
		shared.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   &user,
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	repo := getRepo()
	users, err := repo.GetUsers(context.Background())
	if err != nil {
		shared.Response(c, false, fmt.Sprintf("%v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   &users,
	})
}
