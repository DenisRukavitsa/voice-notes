package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DenisRukavitsa/voice-notes/database"
	"github.com/DenisRukavitsa/voice-notes/modules/user"
	"github.com/DenisRukavitsa/voice-notes/server"
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
	server := server.Create()
	recorder := httptest.NewRecorder()

	body := []byte(`{"email":"test@test.com","password":"secret"}`)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	server.ServeHTTP(recorder, request)

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	log.Println(response)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.NotEmpty(suite.T(), response["userId"])

	userObjectId, err := primitive.ObjectIDFromHex(response["userId"])
	assert.Nil(suite.T(), err)

	filter := bson.D{{Key: "_id", Value: userObjectId}}
	result, err := suite.databaseClient.Database("voice-notes").Collection("users").DeleteOne(context.TODO(), filter)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), int64(1), result.DeletedCount)
}

func (suite *RegisterRouteTestSuite) TestNoPassword() {
	server := server.Create()
	recorder := httptest.NewRecorder()

	body := []byte(`{"email":"test@test.com"}`)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	server.ServeHTTP(recorder, request)

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), "invalid user details", response["error"])
}

func (suite *RegisterRouteTestSuite) TestExistingEmail() {
	collection := suite.databaseClient.Database("voice-notes").Collection("users")
  user := user.UserModel{
		Email: "test@test.com",
		Password: "secret",
	}
	insertResult, _ := collection.InsertOne(context.TODO(), user)
	userObjectId := insertResult.InsertedID.(primitive.ObjectID)
	
	server := server.Create()
	recorder := httptest.NewRecorder()

	body := []byte(`{"email":"test@test.com","password":"secret"}`)
	request, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	server.ServeHTTP(recorder, request)

	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), "user email already registered", response["error"])

	filter := bson.D{{Key: "_id", Value: userObjectId}}
	collection.DeleteOne(context.TODO(), filter)
}

func TestRegisterRouteTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterRouteTestSuite))
}
