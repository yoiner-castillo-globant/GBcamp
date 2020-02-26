package Mongo

import (
    "context"
    "fmt"
    "log"
    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/yoiner-castillo-globant/GBcamp/App/Constants"
)
//poner el error
func GetConn() (*mongo.Client, error) {

clientOptions := options.Client().ApplyURI(Constants.ConnectionString)
// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}
// Check the connection
err = client.Ping(context.TODO(), nil)
if err != nil {
	
	return nil, err
}

fmt.Println("Connected to MongoDB!")

return client, nil

}

func CloseConn(collection *mongo.Collection) error  {

	client := collection.Database().Client()
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB has been Disconnected")
	return err
}
