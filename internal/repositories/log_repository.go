package repositories

import (
	"context"
	"fmt"
	"github.com/nikhildev/gofy/internal/db"
	"github.com/nikhildev/gofy/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository interface {
	GetLog(id string) (*models.LogMessage, error)
	SaveLog(logMessage models.LogMessage) error
}

type logRepository struct {
	store *db.Store
	coll  *mongo.Collection
}

var store, _ = db.NewStore(nil)
var coll = store.Db.Collection("logs")

func (l *logRepository) SaveLog(logMessage models.LogMessage) error {
	_, err := l.coll.InsertOne(nil, logMessage)
	if err != nil {
		return err
	}

	return nil
}

func NewLogRepository() LogRepository {
	return &logRepository{store: store, coll: coll}
}

func (l *logRepository) GetLog(id string) (*models.LogMessage, error) {
	logId, err := primitive.ObjectIDFromHex(id)
	res := coll.FindOne(context.Background(), bson.M{"_id": logId})

	if res.Err() != nil {
		return nil, res.Err()
	}

	doc := &models.LogMessage{}
	err = res.Decode(doc)

	if err != nil {
		fmt.Println("Could not find log message with the id: ", id)
	}

	return doc, nil

}
