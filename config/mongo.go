package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGOURL = "mongodb://localhost:2717"
	Client   *mongo.Client
	Db       *mongo.Database
)

func Init() context.Context {
	ctx := context.TODO()

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGOURL))
	if err != nil {
		log.Fatal(err)
	}

	Db = Client.Database("test")
	// Cats := Db.Collection("cats")

	// result, err := Cats.InsertOne(ctx, bson.D{
	// 	{Key: "çamur", Value: "kahverengi bir kedi"},
	// 	{Key: "breed", Value: "van kedisi"},
	// })

	// fmt.Println(result)
	return ctx
}
