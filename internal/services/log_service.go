package services

import (
	"time"

	"github.com/nikhildev/gofy/internal/models"
	"github.com/nikhildev/gofy/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var logRepo = repositories.NewLogRepository

func CreateLog(logMessage string) error {
	newLogMessage := models.LogMessage{
		ID:        primitive.NewObjectID(),
		Message:   logMessage,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	err := logRepo().SaveLog(newLogMessage)
	if err != nil {
		return err
	}

	return nil
}

func GetLog(id string) (*models.LogMessage, error) {
	logMessage, err := logRepo().GetLog(id)
	if err != nil {
		return nil, err
	}

	return logMessage, nil
}
