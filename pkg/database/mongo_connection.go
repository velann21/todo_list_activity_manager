package database

import (
	"context"
	"github.com/velann21/todo_list_activity_manager/pkg/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var connection *MongoConnection

type MongoConnection struct {
	client *mongo.Client
}

func (mongoConnectionStruct *MongoConnection) MongoConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+helpers.ReadEnv(helpers.MONGOCONNECTIONSTRING)))
	if err != nil {
		return err
	}
	mongoConnectionStruct.client = client
	connection = mongoConnectionStruct
	return nil
}

func GetMongoConnection() *mongo.Client {
	return connection.client
}

func MakeDatabaseAndTodoActivityCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("todo_list1").Collection("todo_activity")
	return collection
}

func MakeDatabaseAndUserTodoCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("todo_list1").Collection("todo_user_activity")
	return collection
}
