package services

import (
	"context"
	"log"

	"github.com/nikhildev/gofy/internal/db/mongodb"
	"github.com/nikhildev/gofy/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getIssuesCollection() *mongo.Collection {
	client, _ := mongodb.NewDbStore(nil, "gofy")
	return mongodb.GetCollection(*client, "issues")
}

var coll = getIssuesCollection()

func CreateIssue(issue models.Issue) (string, error) {

	newIssue, err := coll.InsertOne(context.TODO(), issue)
	if err != nil {
		log.Printf("Error creating issue: %v", err)
	}

	var result models.Issue

	filter := bson.M{"_id": newIssue.InsertedID}
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No issue found with the ID: %v", newIssue.InsertedID)
		} else {
			log.Println(err)
		}
	}

	return result.ID.Hex(), nil
}
