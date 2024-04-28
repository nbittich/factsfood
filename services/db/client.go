package db

import (
	"context"
	"fmt"

	"github.com/nbittich/factsfood/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init() {
	println("connecting to mongo db...")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(fmt.Errorf("could not create mongo client:\n %s", err.Error()))
	}
	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("could not ping mongo:\n %s", err.Error()))
	}
	fmt.Printf("connected!")
	DB = client
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	if err := DB.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func getCollection(collectionName string) *mongo.Collection {
	db := DB.Database(config.MongoDBName, &options.DatabaseOptions{})
	return db.Collection(collectionName, &options.CollectionOptions{})
}
