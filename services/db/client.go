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

var DB *mongo.Client

func init() {
	fmt.Println("connecting to mongo db...")
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

func GetCollection(collectionName string) *mongo.Collection {
	db := DB.Database(config.MongoDBName, &options.DatabaseOptions{})
	return db.Collection(collectionName, &options.CollectionOptions{})
}

func FilterByID(id string) primitive.M {
	return bson.M{"_id": id}
}

func Exist(ctx context.Context, filter bson.M, collection *mongo.Collection) (bool, error) {
	res := collection.FindOne(ctx, filter, &options.FindOneOptions{})

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		} else {
			fmt.Printf("could not call exists, maybe a bug? %s", err)
			return false, err
		}
	}

	return true, nil
}

func FindOneBy[T any](ctx context.Context, filter bson.M, collection *mongo.Collection) (T, error) {
	ptr := new(T)
	res := collection.FindOne(ctx, filter, &options.FindOneOptions{})
	err := res.Decode(ptr)
	return *ptr, err
}

func FindOneByID[T any](ctx context.Context, collection *mongo.Collection, id string) (T, error) {
	return FindOneBy[T](ctx, FilterByID(id), collection)
}

func InsertOrUpdate(ctx context.Context, entity types.Identifiable, collection *mongo.Collection) (string, error) {
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
