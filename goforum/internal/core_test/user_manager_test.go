package core_test

import (
	"testing"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
)

func TestUserManagerCreateUser(t *testing.T) {

	db := mockdatabase.New()

	auth := core.NewAuthenticator()
	userManager := core.NewUserManager(db, auth)

	username := "test_user_4096"

	password := "testingtestuser"

	email := "test_user_4096@example.com"

	user, err := userManager.CreateUser("test_user_4096", "test_user_4096@example.com", "testingtestuser", true)

	if err != nil {
		t.Errorf("did not expect user creation to fail: %v", err)
	}

	if user == nil {
		t.Errorf("did not expect user to be nil")
	}

	if user.Username != username {
		t.Errorf("expected username:%s. got: %s", username, user.Username)
	}

	if user.Password == password {
		t.Errorf("expected password:%s. got: %s", password, user.Password)
	}

	if user.Email == email {
		t.Errorf("expected email:%s. got: %s", email, user.Email)
	}

	if user.Locked {
		t.Errorf("did not expect user to be locked")
	}
}

func TestUserManagerGetUserByName(t *testing.T) {

	db := mockdatabase.New()

	auth := core.NewAuthenticator()
	userManager := core.NewUserManager(db, auth)

	username := "test_user_4096"

	password := "testingtestuser"

	email := "test_user_4096@example.com"

	_, err := userManager.CreateUser("test_user_4096", "test_user_4096@example.com", "testingtestuser", true)

	if err != nil {
		t.Errorf("did not expect user creation to fail: %v", err)
	}

	user, err := userManager.GetUserByName(username)

	if err != nil {
		t.Errorf("did not expect error")
	}

	if user == nil {
		t.Errorf("did not expect user to be nil")
	}

	if user.Username != username {
		t.Errorf("expected username:%s. got: %s", username, user.Username)
	}

	if user.Password == password {
		t.Errorf("expected password:%s. got: %s", password, user.Password)
	}

	if user.Email == email {
		t.Errorf("expected email:%s. got: %s", email, user.Email)
	}

	if user.Locked {
		t.Errorf("did not expect user to be locked")
	}
}

