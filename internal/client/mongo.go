package client

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/BrunoKrugel/go-webhook/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var Collection *mongo.Collection

func InitMongo() (err error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	MongoClient = client

	usersCollection := MongoClient.Database("kiwify-go").Collection("sales")
	Collection = usersCollection
	return nil
}

func UpdateSales(sales model.MongoRequest) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		{Key: "user_id", Value: sales.UserId},
		{Key: "product_id", Value: sales.ProductId},
		{Key: "store_id", Value: sales.StoreId},
		{Key: "date", Value: sales.Date},
	}

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "sales_number", Value: 1}}}}

	_, err = Collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}