package repositories

import (
	"fmt"

	"github.com/nikhildev/gofy/internal/db/mongodb"
	"github.com/nikhildev/gofy/internal/models"
)

type IssueRepository interface {
	CreateIssue(issue *models.Issue) (string, error)
}

type issueRepository struct {
	store mongodb.DbStore
}

func (i *issueRepository) CreateIssue(issue *models.Issue) (string, error) {
	fmt.Println("Creating issue")

	coll := i.store.Db.Collection("issues")
	res, err := coll.InsertOne(nil, issue)
	fmt.Println(res)

	if err != nil {
		panic(err)
	}

	return res.InsertedID.(string), nil
}

func NewIssueRepository(store mongodb.DbStore) IssueRepository {
	return &issueRepository{store: store}
}
