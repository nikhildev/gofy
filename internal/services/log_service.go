package services

import (
	"github.com/nikhildev/gofy/internal/models"
	"github.com/nikhildev/gofy/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var repo = repositories.NewLogRepository()

func CreateLog(logMessage string) error {
	newLogMessage := models.LogMessage{
		ID:        primitive.NewObjectID(),
		Message:   logMessage,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	err := repo.SaveLog(newLogMessage)
	if err != nil {
		return err
	}

	return nil
}

func GetLog(id string) (*models.LogMessage, error) {
	logMessage, err := repo.GetLog(id)
	if err != nil {
		return nil, err
	}

	return logMessage, nil
}
