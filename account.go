package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Rest of the code will go here
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")
	//ash := Trainer{"Ash", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Pewter City"}
	//
	//insertResult, err := collection.InsertOne(context.TODO(), ash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	//trainers := []interface{}{misty, brock}

	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	filter := bson.D{{"name", "Ash"}}

	//update := bson.D{
	//	{"$inc", bson.D{
	//		{"age", 1},
	//	}},
	//}

	//updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//_, _ = collection.DeleteOne(context.TODO(), filter)
	fmt.Printf("Found a single document: %+v\n", result)
}

