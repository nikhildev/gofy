package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbStore struct {
	Db     *mongo.Database
	client *mongo.Client
}

func GetCollection(store DbStore, collName string) *mongo.Collection {
	return store.Db.Collection(collName)
}

func NewDbStore(opts *options.ClientOptions, dbName string) (*DbStore, error) {
	var dbStore *DbStore

	client, db, err := connectToMongo(opts, dbName)

	if err != nil {
		panic("Cannot connect to db")
	}

	if db != nil && client != nil {
		return &DbStore{client: client, Db: db}, nil
	}
	return dbStore, nil
}

func connectToMongo(opts *options.ClientOptions, dbName string) (*mongo.Client, *mongo.Database, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	opts = options.Client().ApplyURI("mongodb://localhost:27777")

	fmt.Println("Connecting to MongoDB")
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Println("Error connecting to MongoDB", err)
		return nil, nil, err
	}

	fmt.Println("Pinging MongoDB")
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("Error pinging MongoDB", err)
		_ = client.Disconnect(ctx)
		return nil, nil, errors.New(fmt.Sprintf("Cannot connect do db. Error: %v", err))
	}
	fmt.Println("Connected successfully to MongoDB")
	var db = client.Database(dbName)

	return client, db, nil
}
