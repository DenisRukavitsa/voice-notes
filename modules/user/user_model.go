package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID       primitive.ObjectID
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
