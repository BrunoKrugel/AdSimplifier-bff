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

var MongoClientOrigin *mongo.Client
var CollectionOrigin *mongo.Collection

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

	usersCollection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))
	Collection = usersCollection
	return nil
}

func UpdateSales(sales model.MongoRequest) (err error) {

	// ReadSales()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		{Key: "user_id", Value: sales.UserId},
		{Key: "product_id", Value: sales.ProductId},
		{Key: "product_name", Value: sales.ProductName},
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

func InitOriginalMongo() (err error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	MongoClientOrigin = client

	usersCollection := MongoClientOrigin.Database("kiwify").Collection("product_sales")
	CollectionOrigin = usersCollection
	return nil
}

func ReadSales() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := CollectionOrigin.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var sales []model.OldRequest
	if err = cursor.All(ctx, &sales); err != nil {
		log.Fatal(err)
	}
	//close mongo connection
	defer func() {
		if err = MongoClientOrigin.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	for _, sale := range sales {
		log.Println(sale)
		WriteSales(sale)
	}
}

func WriteSales(sales model.OldRequest) (err error) {
	//write sales in mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	date, _ := time.Parse("2006-01-02", sales.Date)
	newSales := model.NewRequest{
		User_id:      "dGVzdA==",
		Product_id:   sales.Product_id,
		Product_name: sales.Product_name,
		Store_id:     sales.Store_id,
		Date:         date,
		Sales_number: sales.Sales_number,
	}

	_, err = Collection.InsertOne(ctx, newSales)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
