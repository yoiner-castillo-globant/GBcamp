package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/bson"
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
	
// Set client options
clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
//clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-wlluq.mongodb.net/test?retryWrites=true&w=majority")

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
ash := Trainer{"Ash", 10, "Pallet Town"}
 misty := Trainer{"Misty", 10, "Cerulean City"}
 brock := Trainer{"Brock", 15, "Pewter City"}

 //Insert documents

 //To insert a single document
insertResult, err := collection.InsertOne(context.TODO(), ash)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Inserted a single document: ", insertResult.InsertedID)

//To insert multiple documents at a time
trainers := []interface{}{misty, brock}

insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

//Update Documents
filter := bson.D{{"name", "Ash"}}

update := bson.D{
    {"$inc", bson.D{
        {"age", 1},
    }},
}
updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)


//Find documents
// create a value into which the result can be decoded
var result Trainer

err = collection.FindOne(context.TODO(), filter).Decode(&result)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found a single document: %+v\n", result)

//To find multiple documents

// Pass these options to the Find method
findOptions := options.Find()
//findOptions.SetLimit(2)

// Here's an array in which you can store the decoded documents
var results []*Trainer

// Passing bson.D{{}} as the filter matches all documents in the collection
cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem Trainer
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }

    results = append(results, &elem)
}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}

// Close the cursor once finished
cur.Close(context.TODO())

fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)



//Delete Documents
//you can delete documents using collection.DeleteOne() or collection.DeleteMany().
//bson.D{{}} as the filter argument
//deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
 deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"name", "Ash"}})
 if err != nil {
     log.Fatal(err)
 }
 fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)


var result3 Trainer

err = collection.FindOne(context.TODO(), bson.D{{}}).Decode(&result3)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found a single document: %+v\n", result3)


//To find multiple documents



// Here's an array in which you can store the decoded documents
var results3 []*Trainer

// Passing bson.D{{}} as the filter matches all documents in the collection
cur2, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur2.Next(context.TODO()) {
    
    // create a value into which the single document can be decoded
    var elem Trainer
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }
 
    results3 = append(results3, &elem)
}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}

// Close the cursor once finished
cur2.Close(context.TODO())

for _, element := range results3 {
    // index is the index where we are
    // element is the element from someSlice for where we are
    fmt.Printf(" element: %+v",element)
}
fmt.Printf("Found multiple documents (array of pointers): %+v\n", results3)

}