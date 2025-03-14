// This is just a test user repository
package repositories

import (
	"errors"

	"github.com/nikhildev/gofy/internal/models"
)

type Db []models.User

var userDb = []models.User{}

type UserRepository interface {
	GetAll() *[]models.User
	GetById(id int) *models.User
	Add(user models.User) (*models.User, error)
}

type userRepository struct {
	db *Db
}

func NewUserRepository() UserRepository {
	db := Db(userDb)
	return &userRepository{db: &db}
}

func (r *userRepository) GetAll() *[]models.User {
	users := []models.User(*r.db)
	return &users
}

func (r *userRepository) GetById(id int) *models.User {
	for _, record := range *r.db {
		if record.ID == id {
			return &record
		}
	}
	return nil
}

func (r *userRepository) Add(user models.User) (*models.User, error) {
	id := user.ID
	existingUser := r.GetById(id)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}
	*r.db = append(*r.db, user)
	return &user, nil
}
