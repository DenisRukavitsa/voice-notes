package models

import (
	"context"

	"github.com/DenisRukavitsa/voice-notes/auth"
	"github.com/DenisRukavitsa/voice-notes/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	passwordHash, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	
	collection := database.Database.Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
