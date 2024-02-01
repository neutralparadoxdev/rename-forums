package postgresdb

import (
	"testing"
)

func TestForumRepositoryGetByName(t *testing.T) {
	db, err := New("postgres://goforum:goforum@localhost:5432/goforum_test");

	if err != nil {
		t.Errorf("error creating PostgresDB instance :%v", err)
		return
	}

	defer db.Close()

	repo := db.GetForumRepository()

	forum, err := repo.GetByName("forum_test")

	if err != nil {
		t.Errorf("error getting forum_test:%v", err)
		return
	}

	if forum == nil {
		t.Error("forum was not returned and no error found")
		return
	}

	if !forum.IsPublic {
		t.Error("forum_test was expected to be public")
		return
	}

	if forum.Title != "forum_test" {
		t.Error("forum_test title was expected to be forum_test")
		return
	}

	if forum.Description != "this is a test for the forum" {
		t.Error("forum_test Description was different than expected")
		return
	}
}



