package core_test

import (
	"testing"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
)

func TestSessionManager(t *testing.T) {

	db := mockdatabase.New()

	secret := "SOMETHINGSECRET"

	auth := core.NewAuthenticator()

	userManager := core.NewUserManager(db, auth)

	sessionManager := core.NewSessionManager(
		secret,
		userManager,
		auth,
		db.GetSessionRepository(),
	)

	//func (man *SessionManager) CreateSession(username string, password string) (Session, error) {


	session, err := sessionManager.CreateSession("hello", "hello")

	if err != nil {
		t.Errorf("did not expect an error from CreateSession")
	}

	empty := core.Session{}


	if session == empty {
		t.Errorf("expected session to be not nil")
	}

	if session.Session == "" {
		t.Errorf("expected session Token to not be empty string")
	}

}
