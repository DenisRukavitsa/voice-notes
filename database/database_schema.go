package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func buildDatabaseSchema() error {
	err := createEmailUniqueIndex()
	if err != nil {
		return err
	}
	return nil
}

func createEmailUniqueIndex() error {
	emailUniqueIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	
	usersCollection := Database.Collection("users")
	_, err := usersCollection.Indexes().CreateOne(context.TODO(), emailUniqueIndex)
	if err != nil {
		return err
	}
	return nil
}

