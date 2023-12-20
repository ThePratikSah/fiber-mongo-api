package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
We added omitempty and validate:”required” to the struct
tag to tell Fiber to ignore empty fields and make the field required, respectively.
*/
type User struct {
	Id            primitive.ObjectID `json:"_id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Email         string             `json:"email,omitempty" validate:"required"`
	EmailVerified bool               `json:"emailVerified,omitempty" validate:"required"`
	ImageUrl      string             `json:"imageUrl,omitempty" validate:"required"`
}
