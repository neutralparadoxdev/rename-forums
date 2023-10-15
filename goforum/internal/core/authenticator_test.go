package core

import "testing"

func TestCheck(t *testing.T) {
	auth := NewAuthenticator()

	got, err := auth.Check("hello", "$argon2id$v=19$m=65536,t=3,p=4$9SuvMHDGhFZrZYUPOUvjlQ$4/fDdP7ar1ehn2Bu0Yw1QP0FVUyy33IEHOO8EwctRH8")

	want := true

	if err != nil {
		t.Errorf("got err when none expected %s", err)
	}

	if got != want {
		t.Errorf("got %t when none expected %t", got, want)
	}

}
