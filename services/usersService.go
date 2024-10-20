package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"library.net/module/lib"
	"library.net/module/models"
)

func UserDatabaseConnect() (*DB, error) {
	db, err := ConnectServiceMongo(lib.UriMongo)
	if err != nil {
		fmt.Println("Error Connect: ", err.Error())
		return nil, err
	}
	return db, nil
}

func (db *DB) GetUser(id string) (*models.User, error) {
	user := models.User{}

	//prova gitgit

	collection := db.Client.Database(lib.DatabaseName).Collection("user")

	filter := bson.D{{Key: "_id", Value: id}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error Find: ", err.Error())
		return nil, err
	}

	if err := cur.All(context.TODO(), &user); err != nil {
		fmt.Println("Error Decode: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetAllUser() ([]models.User, error) {
	allUsers := []models.User{}

	collection := db.Client.Database(lib.DatabaseName).Collection("user")

	filter := bson.D{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error Find: ", err.Error())
		return allUsers, err
	}

	if err := cur.All(context.TODO(), &allUsers); err != nil {
		fmt.Println("Error Decode: ", err.Error())
		return allUsers, err
	}

	return allUsers, nil
}
