package Ado

import (
    "fmt"
	"log"
	"errors"
	"context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/yoiner-castillo-globant/GBcamp/App/Mongo"
	"github.com/yoiner-castillo-globant/GBcamp/App/Request"
	"github.com/yoiner-castillo-globant/GBcamp/App/Constants"
	"github.com/yoiner-castillo-globant/GBcamp/App/ApiRest/Logic"
)

func Init() *mongo.Collection {
	
	client, _:= Mongo.GetConn()
	collection:= client.Database(Constants.DBNameMongo).Collection(Constants.NameCollection)
	return collection
}

func AddItemCart(keyCart string, element Request.ArticleCart) error {
	
	collection :=	Init()
	defer Mongo.CloseConn(collection)
	//Update Documents
	filter := bson.D{{"idcart",keyCart}}
	PushToArray := bson.M{"$push": bson.M{"elements": element}}

updateResult, err := collection.UpdateOne(context.TODO(), filter, PushToArray)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}


func ChangeItemAmount(keyCart, keyitem string, amount int) error {
	
	collection :=	Init()
	defer Mongo.CloseConn(collection)
	//Update Documents

	filter := bson.D{
		{"$and", bson.D{
			{"idcart", keyCart},
			{"elements.article.id", keyitem},
		}},
	}

	update := bson.M{"$set": bson.M{"elements.quantity": amount}}

updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func ValidateExistCart(key string) bool  {

	result, err := GetCart(key)
	if err != nil{
		return false
	}

	if result.IdCart == ""{
		return false
	}
	return true
}

func Find(keyCart, keyitem string) Request.Article{
	collection :=	Init()
	defer Mongo.CloseConn(collection)

	var result Request.Article
	filter := bson.D{
		{"$and", bson.D{
			{"idcart", keyCart},
			{"elements.article.id", keyitem},
		}},
	}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil{
		return Request.Article{}
	}

	fmt.Println(result)
	if result.ArticleId == ""{
		return Request.Article{}
	}
	return result

}

func GetCart(key string) (Logic.Cart, error)  {
	collection :=	Init()
	defer Mongo.CloseConn(collection)
	var result Logic.Cart

	filter := bson.D{{"idcart", key}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&result); err != nil{
		return Logic.Cart{}, err
	}

	if result.IdCart == ""{
		return Logic.Cart{}, errors.New("error: The cart is empty")
	}
	return result, nil
}

func InsertCart( datos Logic.Cart ) (bool, error) {

	 if ValidateExistCart(datos.IdCart) {
		return false, errors.New("error: The IdCart is already registered")
	}

collection :=	Init()
defer Mongo.CloseConn(collection)

//To insert a single document
if _, err := collection.InsertOne(context.TODO(), datos); err!= nil{
	return false, err
}
fmt.Println("Inserted document")

	return true, nil
}

func GetCarts() []Logic.Cart{

collection :=	Init()
defer Mongo.CloseConn(collection)
	// Pass these options to the Find method
findOptions := options.Find()
//findOptions.SetLimit(2)

// Here's an array in which you can store the decoded documents
results := []Logic.Cart{}

// Passing bson.D{{}} as the filter matches all documents in the collection
cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {
    // create a value into which the single document can be decoded
    var elem Logic.Cart
	if err := cur.Decode(&elem); err != nil{
		fmt.Println("Error")
	}
	results = append(results, elem)

}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}
// Close the cursor once finished
cur.Close(context.TODO())
return results
}


func GetItemsCarts(keyCart string) ([]*Request.ArticleCart, error){	
	response, err := GetCart(keyCart); 
	if err!= nil{
		return nil, err
	}
	return response.GetAllItems(), nil
}