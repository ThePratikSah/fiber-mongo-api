package controllers

import (
	"context"
	"fiber-mongo-api/src/configs"
	"fiber-mongo-api/src/models"
	"fiber-mongo-api/src/responses"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection = configs.GetCollection(configs.DB, configs.UsersCollection)

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	newUser := models.User{
		Id:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}

	result, err := UsersCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.UserResponse{
		Status:  fiber.StatusCreated,
		Message: "Success",
		Data:    &fiber.Map{"user": result},
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error",
				Data:    &fiber.Map{"error": err.Error()},
			})
		}
		users = append(users, singleUser)
	}

	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"users": users},
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
