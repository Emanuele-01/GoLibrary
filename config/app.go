package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"library.net/module/lib"
)

var App *services

type services struct {
	D *DATABASE
}

type DATABASE struct {
	dbName  string
	dbUrl   string
	timeOut time.Duration
	client  *mongo.Client
}

func NewDatabase(db string, url string) *DATABASE {
	duration := 60
	return &DATABASE{dbName: db, dbUrl: url, timeOut: time.Duration(duration)}
}

func (db *DATABASE) ConnectMDB(minPoolSize uint64, maxPoolSize uint64, maxConnTime uint64) error {
	ctx := context.Background()

	options.Client().SetMinPoolSize(minPoolSize)
	options.Client().SetMaxPoolSize(maxPoolSize)
	options.Client().SetMaxConnIdleTime(time.Duration(maxConnTime) * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.dbUrl))
	if err != nil {
		log.Println("error conn DB: ", err.Error())
		log.Fatal(err)
		return err
	}

	fmt.Println("Conn DB successful...")

	db.client = client

	return nil
}

func (db *DATABASE) GetCollection(collectionName string) *mongo.Collection {
	collection := db.client.Database(lib.DatabaseName).Collection(collectionName)
	return collection
}
