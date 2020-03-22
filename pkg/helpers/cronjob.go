package helpers

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CreateCronJob(cron *cron.Cron, cronPattern string) error {
	_, err := cron.AddFunc(cronPattern, func() {
		DueDateRangeQuery(context.Background())
	})

	if err != nil {
		return err
	}
	go cron.Run()
	return nil
}

func DueDateRangeQuery(ctx context.Context) error {
	timeNow := time.Now()
	formattedCurrentTime := timeNow.UTC().Format(time.RFC3339)
	after := timeNow.Add(5 * time.Minute).UTC().Format(time.RFC3339)
	fmt.Println(formattedCurrentTime)
	fmt.Println(after)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+ReadEnv(MONGOCONNECTIONSTRING)))
	if err != nil {
		return err
	}
	collection := client.Database("todo_list1").Collection("todo_user_activity")

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
