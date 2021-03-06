package dao

import (
	"context"
	"log"
	"time"

	"github.com/LucasFrezarini/go-auth-manager/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var db *mongo.Database

// UserCollection defines the name of the user collection
const UserCollection = "users"

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(env.Config.MongoURI))
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}

	db = client.Database("auth_manager")

	createIndexes()
}

func createIndexes() {
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := db.Collection(UserCollection).Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key:   "email",
			Value: bsonx.Int32(1),
		}},
		Options: options.Index().SetUnique(true),
	}, opts)

	if err != nil {
		log.Panicf("Error while creating indexes on database: %v", err)
	}
}
