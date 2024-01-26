package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
)

type UserRepository struct {
	db *PostgresDatabase
}

func NewUserRepository(db *PostgresDatabase) core.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetByName(username string) (*core.User, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *UserRepository) GetByEmail(email string) (*core.User, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *UserRepository) GetById(id int64) (*core.User, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *UserRepository) Delete(user core.User) error {
	return errors.New("Not Implemented")
}

func (repo *UserRepository) Create(user core.User) (*core.User, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *UserRepository) Save(user core.User) error {
	return errors.New("Not Implemented")
}

func (repo *UserRepository) GetUserNamesForIds(ids []int64) map[int64]string {
	return nil
}


