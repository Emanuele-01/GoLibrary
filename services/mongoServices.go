package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

func ConnectServiceMongo(uri string) (*DB, error) {
	clientOption := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println("Connection services Mongo error: ", err.Error())
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Ping Mongo error: ", err.Error())
		return nil, err
	}

	fmt.Println("Connect to Mongo!!")

	return &DB{Client: client}, nil
}
