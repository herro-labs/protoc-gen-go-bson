package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "example/protogen"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user := &pb.User{
		Name: "John Doe",
		Age:  32,
		Tags: []string{"dev", "employed"},
		Addresses: []*pb.Address{
			{
				Street:  "123 4th Street",
				City:    "Some city",
				Country: "Some country",
			},
		},
		Status:             pb.UserStatus_ACTIVE,
		LastSeenTimestamps: []int64{time.Now().Unix()},
	}
	collection := client.Database("example").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Added user ID:", insertResult.InsertedID)

	filter := bson.D{{Key: "name", Value: "John Doe"}}
	foundUser := pb.UserBSON{}
	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Found user: ", "_id", foundUser.Id, "name", foundUser.Name, "age", foundUser.Age)
}
