package configs

import (
	"context"
	"fiber-mongo-api/src/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = GetCollection(DB, UsersCollection)

func SeedData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users := []interface{}{
		models.User{Name: "John Doe", Email: "johndoe@example.com"},
		models.User{Name: "Jonny Singh", Email: "jonnysingh@example.com"},
	}

	result, err := UserCollection.InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Seeded %d users", len(result.InsertedIDs))
}
