package services

import (
	"context"
	"fiber-mongo-api/src/configs"
	"fiber-mongo-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection = configs.GetCollection(configs.DB, configs.UsersCollection)

func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	result, err := UsersCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer result.Close(ctx)

	for result.Next(ctx) {
		var singleUser models.User
		if err = result.Decode(&singleUser); err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	return users, nil
}

func CreateNewUser(user models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := UsersCollection.InsertOne(ctx, user)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
