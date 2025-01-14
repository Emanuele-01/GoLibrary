package services

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"library.net/module/config"
	"library.net/module/models"
)

type UserService struct {
	conn *config.DATABASE
}

func NewUsrService(connDB *config.DATABASE) *UserService {
	return &UserService{conn: connDB}
}

// The `func (db *DB) GetUser(id string) (*models.User, error) {` function is a method defined on the
// `DB` struct. It is used to retrieve a single user from the database based on the provided `id`.
func (db *UserService) GetUser(id string) (*models.User, error) {

	user := models.User{}
	collection := db.conn.GetCollection("")

	idUser, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error parse ObjectID: ", err.Error())
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: idUser}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error Find: ", err.Error())
		return nil, err
	}

	if err := cur.All(context.TODO(), &user); err != nil {
		log.Println("Error Decode: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (db *UserService) GetAllUser() ([]models.User, error) {
	allUsers := []models.User{}

	collection := db.conn.GetCollection("")

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

func (db *UserService) CreateUser(user models.User, collectionName string) error {
	collection := db.conn.GetCollection(collectionName)

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("error user insert: ", err.Error())
		return err
	}

	return nil
}
