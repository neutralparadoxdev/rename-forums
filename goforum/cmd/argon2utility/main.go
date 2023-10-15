package main

import (
	"fmt"
	"os"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the string to be hashed")
		return
	}

	authenticator := core.NewAuthenticator()
	val, err := authenticator.Generate(os.Args[1])
	if err == nil {
		fmt.Printf("provided: \"%s\" \nvalue: \"%s\"\n", os.Args[1], val)
	} else {
		fmt.Printf("error: %v\n", err)
	}

}
