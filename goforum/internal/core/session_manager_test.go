package core

import (
	"testing"
)

func TestSessionManager(t *testing.T) {
	secret := "AbsoluteSecret"

	userMan := NewUserManager(nil)
	auth := NewAuthenticator()

	got := NewSessionManager(secret, userMan, auth, nil)

	if got.auth != auth {
		t.Errorf("Auth provided in test not the same in object")
	}

	if got.userManager != userMan {
		t.Errorf("User Manager provided in test not the same in object")
	}

	if got.secret != secret {
		t.Errorf("Expected %s Got %s", secret, got.secret)
	}
}
