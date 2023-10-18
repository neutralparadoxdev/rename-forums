package core

import (
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

var (
	CoreErrorAuthenticatorHashingError = NewCoreError(500, "authenticator: failed to hash")
)

func (auth *Authenticator) Generate(rawString string) (string, *CoreError) {
	encoded, err := auth.config.HashEncoded([]byte(rawString))
	if err != nil {
		return "", &CoreErrorAuthenticatorHashingError
	}
	return string(encoded), nil
}

var (
	CoreErrorAuthenticatorFailedToVerify = NewCoreError(501, "authenticator: failed to verify")
)

func (auth *Authenticator) Check(rawString string, hashedString string) (bool, *CoreError) {
	ok, err := argon2.VerifyEncoded([]byte(rawString), []byte(hashedString))

	if err != nil {
		return false, &CoreErrorAuthenticatorFailedToVerify
	}

	return ok, nil
}
