package mockdatabase

import (
	"math/rand"
	"time"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type UserRepository struct {
	users map[string]core.User
}

func NewUserRepository() *UserRepository {

	users := make(map[string]core.User)

	users["hello"] = core.User{
		Username: "hello",
		/// password is hello
		Password: "$argon2id$v=19$m=65536,t=3,p=4$9SuvMHDGhFZrZYUPOUvjlQ$4/fDdP7ar1ehn2Bu0Yw1QP0FVUyy33IEHOO8EwctRH8",
		/// email is : hello@example.com
		Email:        "$argon2id$v=19$m=65536,t=3,p=4$Ng66YSm0nj0sSqiEpJjDhg$S5P4xf2+Ma+QzGISAtdGJEeVaKTFdYdIz+0Dk74xIYY",
		UserId:       12,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		LastLogin:    time.Now(),
		Locked:       false,
	}

	return &UserRepository{
		users: users,
	}
}

func (repo *UserRepository) GetByName(username string) (*core.User, error) {

	value, exists := repo.users[username]

	if exists {
		return &value, nil
	} else {
		return nil, nil
	}
}

func (repo *UserRepository) GetByEmail(email string) (*core.User, error) {

	for _, val := range repo.users {
		if val.Email == email {
			return &val, nil
		}
	}
	return nil, nil
}

func (repo *UserRepository) GetById(id int64) (*core.User, error) {
	for _, val := range repo.users {
		if val.UserId == id {
			return &val, nil
		}
	}
	return nil, nil
}

func (repo *UserRepository) Delete(user core.User) error {
	return nil
}

func (repo *UserRepository) Create(user core.User) (*core.User, error) {

	user.CreatedAt = time.Now()
	user.LastLogin = time.Now()
	user.LastModified = time.Now()
	user.UserId = rand.Int63()
	user.Locked = false

	repo.users[user.Username] = user

	return &user, nil
}

func (repo *UserRepository) Save(user core.User) error {
	return nil
}
