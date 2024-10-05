package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var Client *mongo.Client

func ConnectToMongoDB() {
		errGD := godotenv.Load()
		if errGD != nil {
			log.Fatal(".env file couldn't be loaded")
		}

		uri := os.Getenv("URI")

    ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
    defer cancel()

    var err error
    Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Could not ping MongoDB:", err)
    }

    fmt.Println("Connected to MongoDB!")
}
