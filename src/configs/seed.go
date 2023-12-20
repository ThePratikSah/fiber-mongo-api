package configs

import (
	"context"
	"fiber-mongo-api/src/models"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = GetCollection(DB, UsersCollection)

func SeedData() {
	flagFilePath := "/data/seed_data_initialized.txt"

	if _, err := os.Stat(flagFilePath); err == nil {
		log.Println("Data already seeded, skipping.")
		return
	}

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
	os.Create(flagFilePath)
}
