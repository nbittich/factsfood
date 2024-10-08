package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services/utils"
	"github.com/nbittich/factsfood/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	DB          *mongo.Database
)

type PageOptions struct {
	PageNumber int64                `json:"pageNumber" form:"pageNumber" query:"pageNumber" validate:"required,min=1"`
	PageSize   int64                `json:"pageSize"   form:"pageSize"   query:"pageSize"   validate:"required,min=1"`
	Sort       string               `json:"sort"       form:"sort"       query:"sort" `
	Direction  SortDirection        `json:"direction"  form:"direction"  query:"direction"  validate:"oneof=0 1 -1"`
	MongoOpts  *options.FindOptions `json:"mongoOpts,omitempty" form:"mongoOpts,omitempty" query:"mongoOpts,omitempty"`
}

type SortDirection int8

const (
	DESC SortDirection = -1
	ASC  SortDirection = 1
)

func init() {
	log.Println("connecting to mongo db...")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetMaxPoolSize(uint64(config.MongoMaxConnectionPool)).ApplyURI(mongoURI))
	if err != nil {
		panic(fmt.Errorf("could not create mongo client:\n %s", err.Error()))
	}
	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("could not ping mongo:\n %s", err.Error()))
	}
	log.Printf("connected!")
	mongoClient = client
	DB = mongoClient.Database(config.MongoDBName, &options.DatabaseOptions{})
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	if err := mongoClient.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName, &options.CollectionOptions{})
}

func FilterByID(id string) primitive.M {
	return bson.M{"_id": id}
}

func Exist(ctx context.Context, filter bson.M, collection *mongo.Collection) (bool, error) {
	c, err := Count(ctx, filter, collection)
	if err != nil {
		log.Printf("could not call exists, maybe a bug? %s \n", err)
		return false, err
	}
	return c != 0, nil
}

func FindOneBy[T types.HasID](ctx context.Context, filter bson.M, collection *mongo.Collection) (T, error) {
	ptr := new(T)
	res := collection.FindOne(ctx, filter, &options.FindOneOptions{})
	err := res.Decode(ptr)
	return *ptr, err
}

func Count(ctx context.Context, filter interface{}, collection *mongo.Collection) (int64, error) {
	return collection.CountDocuments(ctx, filter)
}

func CountAll(ctx context.Context, collection *mongo.Collection) (int64, error) {
	return Count(ctx, &bson.D{}, collection)
}

func Find[T types.HasID](ctx context.Context, filter interface{}, collection *mongo.Collection, page *PageOptions) ([]T, error) {
	opts := &options.FindOptions{}
	resultSize := 100
	if page != nil {
		if page.MongoOpts != nil {
			opts = page.MongoOpts
		} else {
			if err := utils.ValidateStruct(page); err != nil {
				return nil, err
			}
			skip := (page.PageNumber - 1) * page.PageSize
			opts.SetSkip(skip)
			opts.SetLimit(page.PageSize)
			resultSize = int(page.PageSize)
			if page.Sort != "" {
				opts.SetSort(bson.M{page.Sort: page.Direction})
			}

		}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return CursorToSlice[T](ctx, cursor, resultSize)
}

func CursorToSlice[T types.HasID](ctx context.Context, cursor *mongo.Cursor, size int) ([]T, error) {
	defer cursor.Close(ctx)

	results := make([]T, 0, size)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func FindAll[T types.HasID](ctx context.Context, collection *mongo.Collection, page *PageOptions) ([]T, error) {
	return Find[T](ctx, &bson.D{}, collection, page)
}

func FindOneByID[T types.HasID](ctx context.Context, collection *mongo.Collection, id string) (T, error) {
	return FindOneBy[T](ctx, FilterByID(id), collection)
}

func Save[T types.Identifiable](entity T, col *mongo.Collection) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	return InsertOrUpdate(ctx, entity, col)
}

func InsertOrUpdateMany(ctx context.Context, entities []types.Identifiable, collection *mongo.Collection) error {
	models := make([]mongo.WriteModel, 0, len(entities))
	for _, entity := range entities {
		if entity.GetID() == "" {
			entity.SetID(uuid.New().String())
			models = append(models, mongo.NewInsertOneModel().SetDocument(entity))
		} else {
			models = append(models, mongo.NewReplaceOneModel().SetUpsert(true).SetReplacement(entity).SetFilter(FilterByID(entity.GetID())))
		}
	}
	_, err := collection.BulkWrite(ctx, models, &options.BulkWriteOptions{})
	return err
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
		option := &options.ReplaceOptions{}
		option.SetUpsert(true)
		_, err = collection.ReplaceOne(ctx, FilterByID(entity.GetID()), entity, option)
	}
	return id, err
}
