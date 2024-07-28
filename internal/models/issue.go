package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Issue struct {
	Assignee    string
	Cc          []string
	Component   string
	CreatedAt   time.Time
	Description string
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted   bool
	Priority    IssuePriority
	Reporter    string
	Severity    IssueSeverity
	Status      IssueStatus
	Title       string
	UpdatedAt   time.Time
	ChangeLog   []Change
}

type IssueStatus string
type IssuePriority string
type IssueSeverity string

const (
	OPEN        IssueStatus = "OPEN"
	CLOSED      IssueStatus = "CLOSED"
	IN_PROGRESS IssueStatus = "IN_PROGRESS"
)
