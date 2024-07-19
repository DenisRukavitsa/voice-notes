package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/modules/user"
	"github.com/DenisRukavitsa/voice-notes/tests/helpers"
	"github.com/gofor-little/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterRouteTestSuite struct {
	suite.Suite
	databaseClient *mongo.Client
}

func (suite *RegisterRouteTestSuite) SetupSuite() {
	if err := env.Load("../.env"); err != nil {
		panic(err)
	}
	suite.databaseClient = database.Connect()
}

func (suite *RegisterRouteTestSuite) TearDownSuite() {
	database.Disconnect(suite.databaseClient)
}

func (suite *RegisterRouteTestSuite) TestSuccess() {
	body := []byte(`{"email":"test@test.com","password":"secret"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/register", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, code)
	assert.NotEmpty(suite.T(), response["userId"])

	userObjectId, err := primitive.ObjectIDFromHex(response["userId"])
	assert.Nil(suite.T(), err)
	filter := bson.D{{Key: "_id", Value: userObjectId}}
	result, err := suite.databaseClient.Database("voice-notes").Collection("users").DeleteOne(context.TODO(), filter)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result.DeletedCount)
}

func (suite *RegisterRouteTestSuite) TestNoPassword() {
	body := []byte(`{"email":"test@test.com"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/register", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, code)
	assert.Equal(suite.T(), "invalid user details", response["error"])
}

func (suite *RegisterRouteTestSuite) TestExistingEmail() {
	collection := suite.databaseClient.Database("voice-notes").Collection("users")
  user := user.UserDto{
		Email: "test@test.com",
		Password: "secret",
	}
	insertResult, _ := collection.InsertOne(context.TODO(), user)
	userObjectId := insertResult.InsertedID.(primitive.ObjectID)
	
	body := []byte(`{"email":"test@test.com","password":"secret"}`)
	code, response, err := helpers.SendRequest(http.MethodPost, "/register", body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, code)
	assert.Equal(suite.T(), "user email already registered", response["error"])

	filter := bson.D{{Key: "_id", Value: userObjectId}}
	collection.DeleteOne(context.TODO(), filter)
}

func TestRegisterRouteTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterRouteTestSuite))
}
