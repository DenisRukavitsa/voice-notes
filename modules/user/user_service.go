package user

import (
	"context"
	"log"

	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/modules/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func save(user UserDto) (primitive.ObjectID, error) {
	log.Println("saving user")
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

func findUserByEmail(email string) (UserModel, error) {
	var user UserModel
	collection := database.Database.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}

func checkUserCredentials(user UserDto) (string, error) {
	storedUser, err := findUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	isPasswordValid := auth.CheckPasswordHash(user.Password, storedUser.Password)
	if !isPasswordValid {
		return "", nil
	}

	return storedUser.Id, nil
}
