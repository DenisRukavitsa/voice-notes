package user

import (
	"context"
	"log"

	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/modules/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func save(user UserModel) (primitive.ObjectID, error) {
	log.Println("saving user", user)
	passwordHash, err := auth.HashPassword(user.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = passwordHash
	
	collection := database.Database.Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
