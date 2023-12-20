package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Resume struct {
	Id         primitive.ObjectID `json:"_id,omitempty"`
	UniqueText string             `json:"uniqueText,omitempty"`
	UserId     string             `json:"userId,omitempty"`
}
