package db

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// We create a struct store the database connection instance and the client instance
type Store struct {
	Db     *mongo.Database
	Client *mongo.Client
}

func GetCollection(store Store, collName string) *mongo.Collection {
	return store.Db.Collection(collName)
}

// This needs to be called from a function that has a defer statement to close the connection
func NewStore(opts *options.ClientOptions) (*Store, error) {
	var store *Store

	//The dbName is hardcoded here. It should be passed as a parameter probably from env variables
	dbName := "gofy"

	// If the store and client is not nil, we reuse the connection
	//if &store.Db != nil && &store.client != nil {
	//	return store, nil
	//}

	// If the store and client is nil, we create a new connection
	client, db, err := connectToMongo(opts, dbName)

	if err != nil {
		panic("Cannot connect to db")
	}

	if db != nil && client != nil {
		return &Store{Client: client, Db: db}, nil
	}
	return store, nil
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
