package core

import "time"

type User struct {
	Username     string
	Password     string
	UserId       int64
	Email        string
	CreatedAt    time.Time
	LastModified time.Time
	LastLogin    time.Time
	Locked       bool
}

func DefaultUser() User {
	return User{
		Username:     "Something",
		Password:     "Else",
		Email:        "Somethingelse@example.com",
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		LastLogin:    time.Now(),
		UserId:       121231231,
		Locked:       false,
	}
}
