package core_test

import (
	"testing"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
)

func TestCommentManagerCreateCommentForPost(t *testing.T) {
	db := mockdatabase.New()

	commentManager := core.NewCommentManager(db)

	secret := "SOMETHINGSECRET"

	auth := core.NewAuthenticator()

	userManager := core.NewUserManager(db, auth)

	sessionManager := core.NewSessionManager(
		secret,
		userManager,
		auth,
		db.GetSessionRepository(),
	)

	session, err := sessionManager.CreateSession("hello", "hello")

	if err != nil {
		t.Errorf("did not expect error from session %v", err)
		return
	}

	empty := core.Session{}

	if session == empty {
		t.Errorf("did not expect empty session")
		return
	}

	ok, cerr := sessionManager.VerifySession(&session)

	if cerr != nil {
		t.Errorf("sessionManager Verification failed: %v", cerr)
		return
	}

	if !ok {
		t.Errorf("failed to verify session")
		return
	}

	id := int64(1)

	text := "This is a nice post"

	commentId, err := commentManager.CreateCommentForPost(session, id, text)

	if err != nil {
		t.Errorf("did not expect error from CreateCommentForPost %v", err)
		return
	}

	if commentId == 0 {
		t.Errorf("did not expect commentId to be zero")
	}
}

func TestCommentManagerCreateCommentForComment(t *testing.T) {
	db := mockdatabase.New()

	commentManager := core.NewCommentManager(db)

	secret := "SOMETHINGSECRET"

	auth := core.NewAuthenticator()

	userManager := core.NewUserManager(db, auth)

	sessionManager := core.NewSessionManager(
		secret,
		userManager,
		auth,
		db.GetSessionRepository(),
	)

	session, err := sessionManager.CreateSession("hello", "hello")

	if err != nil {
		t.Errorf("did not expect error from session %v", err)
		return
	}

	empty := core.Session{}

	if session == empty {
		t.Errorf("did not expect empty session")
		return
	}

	ok, cerr := sessionManager.VerifySession(&session)

	if cerr != nil {
		t.Errorf("sessionManager Verification failed: %v", cerr)
		return
	}

	if !ok {
		t.Errorf("failed to verify session")
		return
	}

	id := int64(10)

	commentId, err := commentManager.CreateCommentForComment(session, id, "Test message")

	if err != nil {
		t.Errorf("commentManager.CreateCommentForComment errored: %v", err)
		return
	}

	if commentId == 0 {
		t.Errorf("did not expect commentId to be zero")
	}
}


