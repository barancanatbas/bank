package repository

import (
	"context"
	"log"
	"mongoexample/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoBank interface {
	Create(acount models.Account) error
	Info(objeId string) ([]primitive.M, error)
	AddMoney(filter primitive.M, update primitive.M) error
	ReduceMoney(filter primitive.M, update primitive.M) error
}

type bankRepository struct {
	Ctx        context.Context
	Db         *mongo.Database
	Collection *mongo.Collection
}

// compile time proff
var _ RepoBank = bankRepository{}

func Bank(db *mongo.Database, ctx context.Context) bankRepository {
	coll := db.Collection("bank")
	return bankRepository{
		Db:         db,
		Collection: coll,
		Ctx:        ctx,
	}
}

// for new account
func (b bankRepository) Create(acount models.Account) error {
	_, err := b.Collection.InsertOne(b.Ctx, acount)
	return err
}

// get info for account
func (b bankRepository) Info(objeId string) ([]primitive.M, error) {
	id, _ := primitive.ObjectIDFromHex(objeId)
	filter := bson.M{"_id": id}
	cur, err := b.Collection.Find(b.Ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(b.Ctx) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(b.Ctx)
	return results, err
}

// add money in account
func (b bankRepository) AddMoney(filter primitive.M, update primitive.M) error {
	_, err := b.Collection.UpdateOne(b.Ctx, filter, update)
	return err
}

// reduce money in account
func (b bankRepository) ReduceMoney(filter primitive.M, update primitive.M) error {
	_, err := b.Collection.UpdateOne(b.Ctx, filter, update)
	return err
}
