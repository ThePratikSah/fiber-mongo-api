package controllers

import (
	"fiber-mongo-api/src/configs"
	"fiber-mongo-api/src/models"
	"fiber-mongo-api/src/responses"
	"fiber-mongo-api/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection = configs.GetCollection(configs.DB, configs.UsersCollection)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

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

	result, err := services.CreateNewUser(newUser)
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
	results, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error",
			Data:    &fiber.Map{"error": err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &fiber.Map{"users": results},
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
