package repository

import (
	"context"
	"go-gin-repository/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongoDB struct {
	DB *mongo.Database
}

func NewUserMongoDB(db *mongo.Database) UserRepository {
	return &userMongoDB{DB: db}
}

func (repo *userMongoDB) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user.ObjectID = primitive.NewObjectID()
	user.Id = user.ObjectID.Hex()
	_, err := repo.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *userMongoDB) GetUsers(ctx context.Context) ([]entity.User, error) {
	cur, err := repo.DB.Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []entity.User
	for cur.Next(context.Background()) {
		user := entity.User{} // satuan
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		user.Id = user.ObjectID.Hex()
		users = append(users, user)
	}

	if users == nil {
		users = []entity.User{}
	}
	return users, nil
}
