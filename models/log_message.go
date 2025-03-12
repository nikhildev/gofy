package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Message   string             `bson:"message"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	Type      LogType            `bson:"type"`
}

type LogType string

const (
	CREATE LogType = "CREATE"
	UPDATE LogType = "UPDATE"
	DELETE LogType = "DEL"
)
