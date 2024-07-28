package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikhildev/gofy/internal/db/mongodb"
	"github.com/nikhildev/gofy/internal/models"
	"github.com/nikhildev/gofy/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetIssuesCollection() *mongo.Collection {
	client, _ := mongodb.NewDbStore(nil, "gofy")
	return mongodb.GetCollection(*client, "issues")
}

func CreateIssueHandler(c *fiber.Ctx) error {
	newIssue := &models.Issue{
		Title:       "Test updated",
		Description: "This is the issue description",
		Reporter:    "Test",
		Component:   "Test",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
		Status:      models.IssueStatus("OPEN"),
		Priority:    models.IssuePriority("LOW"),
		Severity:    models.IssueSeverity("LOW"),
	}

	insertedId, err := services.CreateIssue(*newIssue)

	if err != nil {
		fmt.Printf("Error creating issue: %v", err)
	}

	return c.SendString(insertedId)

}
