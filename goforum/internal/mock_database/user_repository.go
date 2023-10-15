package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type UserRepository struct {
}

func (repo *UserRepository) GetByName(username string) (*core.User, error) {
	user := core.User{
		Username: "hello",
		/// password is hello
		Password: "$argon2id$v=19$m=65536,t=3,p=4$9SuvMHDGhFZrZYUPOUvjlQ$4/fDdP7ar1ehn2Bu0Yw1QP0FVUyy33IEHOO8EwctRH8",
		/// email is : hello@example.com
		Email:  "$argon2id$v=19$m=65536,t=3,p=4$Ng66YSm0nj0sSqiEpJjDhg$S5P4xf2+Ma+QzGISAtdGJEeVaKTFdYdIz+0Dk74xIYY",
		Locked: false,
	}

	return &user, nil
}

func (repo *UserRepository) GetById() (*core.User, error) {
	user := core.User{
		Username: "hello",
		/// password is hello
		Password: "$argon2id$v=19$m=65536,t=3,p=4$9SuvMHDGhFZrZYUPOUvjlQ$4/fDdP7ar1ehn2Bu0Yw1QP0FVUyy33IEHOO8EwctRH8",
		/// email is : hello@example.com
		Email:  "$argon2id$v=19$m=65536,t=3,p=4$Ng66YSm0nj0sSqiEpJjDhg$S5P4xf2+Ma+QzGISAtdGJEeVaKTFdYdIz+0Dk74xIYY",
		Locked: false,
	}

	return &user, nil
}

func (repo *UserRepository) Delete(user core.User) error {
	return nil
}

func (repo *UserRepository) Create(user core.User) (*core.User, error) {
	return nil, nil
}

func (repo *UserRepository) Save(user core.User) error {
	return nil
}
