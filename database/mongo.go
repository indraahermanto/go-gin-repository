package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectMongoDB() {
	var (
		DBHost     = os.Getenv("DB_HOST")
		DBPort     = os.Getenv("DB_PORT")
		DBUsername = os.Getenv("DB_USERNAME")
		DBPassword = os.Getenv("DB_PASSWORD")
		DBName     = os.Getenv("DB_DATABASE")
	)

	// connect to the database
	fmt.Println("Connecting to MongoDB")
	dbUrl := "mongodb://" + DBUsername + ":" + DBPassword + "@" + DBHost + ":" + DBPort
	fmt.Println(dbUrl)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	MongoDB = client.Database(DBName)
}
