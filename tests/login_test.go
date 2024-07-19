package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/modules/auth"
	"github.com/DenisRukavitsa/voice-notes/modules/user"
	"github.com/DenisRukavitsa/voice-notes/tests/helpers"
	"github.com/gofor-little/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRouteTestSuite struct {
	suite.Suite
	databaseClient *mongo.Client
	usersCollection *mongo.Collection
	userObjectId primitive.ObjectID
}

func (suite *LoginRouteTestSuite) SetupSuite() {
	if err := env.Load("../.env"); err != nil {
		panic(err)
	}
	suite.databaseClient = database.Connect()
	suite.usersCollection = suite.databaseClient.Database("voice-notes").Collection("users")

	password, _ := auth.HashPassword("secret")
  user := user.UserDto{
		Email: "test@test.com",
		Password: password,
	}
	insertResult, _ := suite.usersCollection.InsertOne(context.TODO(), user)
	suite.userObjectId = insertResult.InsertedID.(primitive.ObjectID)
}

func (suite *LoginRouteTestSuite) TearDownSuite() {
	filter := bson.D{{Key: "_id", Value: suite.userObjectId}}
	suite.usersCollection.DeleteOne(context.TODO(), filter)
	database.Disconnect(suite.databaseClient)
}

func (suite *LoginRouteTestSuite) TestSuccess() {
	body := []byte(`{"email":"test@test.com","password":"secret"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/login", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, code)
	assert.NotEmpty(suite.T(), response["accessToken"])
}

func (suite *LoginRouteTestSuite) TestNoEmail() {
	body := []byte(`{"email":"","password":"secret"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/login", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, code)
	assert.Equal(suite.T(), "invalid user details", response["error"])
}

func (suite *LoginRouteTestSuite) TestWrongPassword() {
	body := []byte(`{"email":"test@test.com","password":"wrong"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/login", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusUnauthorized, code)
	assert.Equal(suite.T(), "invalid user credentials", response["error"])
}

func TestLoginRouteTestSuite(t *testing.T) {
	suite.Run(t, new(LoginRouteTestSuite))
}
