package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB          *mongo.Client
	initialized bool
)

func init() {
	if initialized {
		return
	}
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
	initialized = true
}

func Disconnect() {
	if !initialized {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	if err := DB.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	db := DB.Database(config.MongoDBName, &options.DatabaseOptions{})
	return db.Collection(collectionName, &options.CollectionOptions{})
}

func FilterByID(id string) primitive.M {
	return bson.M{"_id": id}
}

func InsertOrUpdate(entity types.Identifiable, collectionName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	collection := GetCollection(collectionName)
	var err error
	id := entity.GetID()
	if id == "" {
		entity.SetID(uuid.New().String())
		_, err = collection.InsertOne(ctx, entity, &options.InsertOneOptions{})
		if err != nil {
			return "", err
		}
	} else {
		_, err = collection.UpdateOne(ctx, FilterByID(entity.GetID()), entity, &options.UpdateOptions{})
	}
	return id, err
}
