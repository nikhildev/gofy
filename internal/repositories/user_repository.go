package repositories

// This is just a test user repository

type User struct {
	ID   int
	Name string
}

type Db []User

var userDb = []User{}

type UserRepository interface {
	GetAll() *[]User
	GetById(id int) *User
	Add(record User) *User
}

type userRepository struct {
	db *Db
}

func NewUserRepository() UserRepository {
	db := Db(userDb)
	return &userRepository{db: &db}
}

func (r *userRepository) GetAll() *[]User {
	users := []User(*r.db)
	return &users
}

func (r *userRepository) GetById(id int) *User {
	for _, record := range *r.db {
		if record.ID == id {
			return &record
		}
	}
	return nil
}

func (r *userRepository) Add(record User) *User {
	*r.db = append(*r.db, record)
	return &record
}
