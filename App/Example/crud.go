package main

import (
	"context" // manage multiple requests
	"encoding/json" // Use JSON encoding for bson. M string
	"fmt" // Println() function
	"log"
	"reflect" // get an object type
	
	// import 'mongo-go-driver' package libraries
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	)

type MongoFields struct {
	Key string `json:"key,omitempty"`
	// ObjectId() or objectid. ObjectID is deprecated--use primitive instead
	ID primitive.ObjectID `bson:"_id, omitempty"`
	
	// Use these field tags so Golang knows how to map MongoDB fields
	// `bson:"string field" json:"string field"`
	StringField string `bson:"string field" json:"string field"`
	IntField int `bson:"int field" json:"int field"`
	BoolField bool `bson:"bool field" json:"bool field"`
	}

	func main() {

		// Declare host and port options to pass to the Connect() method
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		
		// Connect to the MongoDB and return Client instance
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
		fmt. Println("mongo. Connect() ERROR:", err)
		log. Fatal(err)
		}
		
		// Declare Context object for the MongoDB API calls
		ctx := context.Background()
		
		// Access a MongoDB collection through a database
		col := client.Database("BootCamp").Collection("Cart")


data := MongoFields{Key: "xxx", BoolField: true, IntField: 23,StringField: "otro dato"}
_, err = col.InsertOne(context.TODO(), data)

		filter := bson.M{"int field": bson.M{"$gt":42}}
// Create a string using ` string escape ticks
query := `{"$eq":"last value"}`

// Declare an empty BSON Map object
var bsonMap bson.M
// Use the JSON package's Unmarshal() method
err = json.Unmarshal([]byte(query), &bsonMap)
if err != nil {
log.Fatal("json. Unmarshal() ERROR:", err)
} else {
fmt.Println("bsonMap:", bsonMap)
fmt.Println("bsonMap TYPE:", reflect.TypeOf(bsonMap))
fmt.Println("BSON:", reflect.TypeOf(bson.M{"int field": bson.M{"$gt":42}}))
}
// Nest the Unmarshalled BSON map inside another BSON object
filter = bson.M{"string field": bsonMap}

// Pass the filter to Find() to return a MongoDB cursor
cursor, err := col.Find(ctx, filter)
if err != nil {
log.Fatal("col. Find ERROR:", err)
}

// Print cursor object
fmt.Println("\ncursor TYPE:", reflect.TypeOf(cursor))
fmt.Println("cursor:", cursor)

// iterate through all documents
for cursor.Next(ctx) {
var p MongoFields

// Decode the document
if err := cursor.Decode(&p); err != nil {
log.Fatal("cursor. Decode ERROR:", err)
}
// Print the results of the iterated cursor object
fmt.Printf("\nMongoFields: %+v\n", p)
}
}