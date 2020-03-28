package dao

import (
	"context"
	"fmt"
	"github.com/velann21/todo_list_activity_manager/pkg/database"
	databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TodoDaoActivityImpl struct {
	 MongoConnection *mongo.Client
}

// TODO Make more comments  on all the DAO's
func (todoActivity *TodoDaoActivityImpl) AddTodoActivity(ctx context.Context, req *requests.TodoRequest) (*string, error) {
	collection := database.MakeDatabaseAndTodoActivityCollection(todoActivity.MongoConnection)
	res, err := collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		objectID := oid.Hex()
		return &objectID, nil
	} else {
		return nil, nil
	}
}
func (todoActivity *TodoDaoActivityImpl) GetTodoActivity(ctx context.Context, activityID string) (*requests.TodoRequest, error) {
	collection := database.MakeDatabaseAndTodoActivityCollection(todoActivity.MongoConnection)
	objectid, err := primitive.ObjectIDFromHex(activityID)
	if err != nil {
		return nil, err
	}
	var todoModel requests.TodoRequest
	if err := collection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&todoModel); err != nil {
		return nil, err
	}
	return &todoModel, nil

}
func (todoActivity *TodoDaoActivityImpl) UpdateTodoActivity(ctx context.Context, activityID string, req *requests.UpdateTodoStruct) error {
	collection := database.MakeDatabaseAndTodoActivityCollection(todoActivity.MongoConnection)
	oID, err := primitive.ObjectIDFromHex(activityID)
	if err != nil {
		return err
	}
	fmt.Println(oID)
	_, err = collection.ReplaceOne(ctx, bson.D{{"_id", oID}}, req)
	if err != nil {
		return err
	}
	return nil

}

func (todoActivity *TodoDaoActivityImpl) DeleteTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error {
	collection := database.MakeDatabaseAndTodoActivityCollection(todoActivity.MongoConnection)
	oID, err := primitive.ObjectIDFromHex(req.ActivityID)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oID})
	if err != nil {
		return err
	}
	return nil
}

func (todoActivity *TodoDaoActivityImpl) InsertUserActivity(ctx context.Context, req *databaseModel.UserBasedTaskModel) (*string, error) {
	collection := database.MakeDatabaseAndUserTodoCollection(todoActivity.MongoConnection)
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &req.ID, nil
}

func (todoActivity *TodoDaoActivityImpl) UpdateUSerActivity(ctx context.Context, userID string, task []databaseModel.Task) error {
	collection := database.MakeDatabaseAndUserTodoCollection(todoActivity.MongoConnection)
	update := bson.D{{"$set", bson.D{{"tasks", task}}}}
	_, err := collection.UpdateOne(ctx, bson.D{{"_id", userID}}, update)
	if err != nil {
		return err
	}
	return nil
}

func (todoActivity *TodoDaoActivityImpl) DeleteUserTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error {
	return nil
}

func (todoActivity *TodoDaoActivityImpl) GetUserTodoActivity(ctx context.Context, userID string) (*databaseModel.UserBasedTaskModel, error) {
	collection := database.MakeDatabaseAndUserTodoCollection(todoActivity.MongoConnection)
	var model databaseModel.UserBasedTaskModel
	if err := collection.FindOne(ctx, bson.D{{"_id", userID}}).Decode(&model); err != nil {
		return nil, err
	}
	return &model, nil
}

func (todoActivity *TodoDaoActivityImpl) AddSubTask(ctx context.Context, req []interface{}) (*[]string, error) {
	collection := database.MakeDatabaseAndUserTodoCollection(todoActivity.MongoConnection)
	insertMany, err := collection.InsertMany(ctx, req)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0)
	for i, v := range insertMany.InsertedIDs {
		if oid, ok := v.(primitive.ObjectID); ok {
			objectID := oid.Hex()
			ids[i] = objectID
		}
	}
	return &ids, nil
}

func (todoActivity *TodoDaoActivityImpl) UpdateUserTodoActivity(ctx context.Context, userID string, activityID string) error {
	collection := database.MakeDatabaseAndUserTodoCollection(todoActivity.MongoConnection)
	// TODO make this as an TYPE/CONST
	update := bson.M{"$pull": bson.M{"tasks": bson.M{"taskid": activityID}}}
	updates, err := collection.UpdateOne(ctx, bson.D{{"_id", userID}}, update)
	if err != nil {
		return err
	}
	fmt.Println(updates)
	return nil
}

func (todoActivity *TodoDaoActivityImpl) DueDateRangeQuery(ctx context.Context, timeDiff int) error {
	timeNow := time.Now()
	formattedCurrentTime := timeNow.UTC().Format(time.RFC3339)
	after := timeNow.Add(5 * time.Minute).UTC().Format(time.RFC3339)
	collection := database.MakeDatabaseAndTodoActivityCollection(todoActivity.MongoConnection)
	filter := bson.M{"customduedate": bson.M{"$gt": formattedCurrentTime, "$lt": after}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		fmt.Println("")
	}
	fmt.Println(episodes)
	return nil
}
