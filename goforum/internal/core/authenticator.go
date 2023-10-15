package core

import (
	"errors"

	"github.com/matthewhartstonge/argon2"
)

type Authenticator struct {
	config argon2.Config
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{
		config: argon2.DefaultConfig(),
	}
}

func (auth *Authenticator) Generate(rawString string) (string, error) {
	encoded, err := auth.config.HashEncoded([]byte(rawString))
	if err != nil {
		return "", errors.New("could not encode string")
	}
	return string(encoded), nil
}

func (auth *Authenticator) Check(rawString string, hashedString string) (bool, error) {
	ok, err := argon2.VerifyEncoded([]byte(rawString), []byte(hashedString))

	if err != nil {
		return false, err
	}

	return ok, nil
}
