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
var SalesCollection *mongo.Collection
var InfoCollection *mongo.Collection

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

	SalesCollection = MongoClient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION_SALES"))
	InfoCollection = MongoClient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION_INFO"))
	return nil
}

func UpdateSales(sales model.MongoSales, incValue int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		{Key: "user_id", Value: sales.User_id},
		{Key: "product_id", Value: sales.Product_id},
		{Key: "product_name", Value: sales.Product_name},
		{Key: "store_id", Value: sales.Store_id},
		{Key: "date", Value: sales.Date},
	}

	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "sales_number", Value: incValue}}}}

	_, err = SalesCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func InsertSales(sales model.MongoSalesInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = InfoCollection.InsertOne(ctx, sales)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteSales(sales model.MongoSalesInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "user_id", Value: sales.User_id},
		{Key: "product_name", Value: sales.Product_name},
		{Key: "store_id", Value: sales.Store_id},
		{Key: "order_ref", Value: sales.Order_ref},
	}

	_, err = InfoCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CloseMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient.Disconnect(ctx)
}
