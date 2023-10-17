package core

import (
	"errors"
	"log"
)

type UserManager struct {
	db   Database
	auth *Authenticator
}

func NewUserManager(db Database, auth *Authenticator) *UserManager {
	return &UserManager{
		db:   db,
		auth: auth,
	}
}

func (man *UserManager) CreateUser(username string, email string, password string, acceptedEula bool) (*User, error) {

	if !acceptedEula {
		return nil, errors.New("create_user: must accept eula")
	}

	repo := man.db.GetUserRepository()

	existingUser, err := repo.GetByName(username)

	if err != nil {
		return nil, errors.New("create_user: db error on checking for username")
	}

	if existingUser != nil {
		return nil, errors.New("create_user: user exists")
	}

	hashedEmail, err := man.auth.Generate(email)
	if err != nil {
		return nil, errors.New("create_user: could not hash email")
	}

	existingUser, err = repo.GetByEmail(hashedEmail)
	if err != nil {
		return nil, errors.New("create_user: could not get by email")
	}

	if existingUser != nil {
		log.Printf("email exists: %s(%s)", email, hashedEmail)
		return nil, errors.New("create_user: pre-existing email found")
	}

	hashedPassword, err := man.auth.Generate(password)
	if err != nil {
		return nil, errors.New("create_user: could not hash password")
	}

	user := User{
		Username: username,
		Email:    hashedEmail,
		Password: hashedPassword,
	}

	completedUser, err := repo.Create(user)

	if err != nil {
		return nil, errors.New("create_user: could not create new user")
	} else {
		return completedUser, nil
	}
}

func (man *UserManager) GetUserByName(username string) (*User, error) {
	rep := man.db.GetUserRepository()
	user, err := rep.GetByName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
