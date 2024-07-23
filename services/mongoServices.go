package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectServiceMongo(uri, dbName string) (*MongoService, error) {
	clientOption := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		fmt.Println("Connection services Mongo error: ", err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Ping Mongo error: ", err.Error())
		return nil, err
	}

	fmt.Println("Connect to Mongo!!")

	database := client.Database(dbName)

	return &MongoService{Client: client, Database: database}, nil
}
