package main

import (
	"fmt"
	"github.com/rollout/rox-go/v5/server"
)

// Create Rox flags in the Flags container class
type Flags struct {
	EnableTutorial server.RoxFlag
}

var flags = &Flags{
	// Define the feature flags
	EnableTutorial: server.NewRoxFlag(false),
}

var rox *server.Rox

func main() {
	options := server.NewRoxOptions(server.RoxOptionsBuilder{})

	rox := server.NewRox()

	// Register the flags container with the CloudBees platform
	rox.RegisterWithEmptyNamespace(flags)

	// Setup the feature management environment key
	<-rox.Setup("340ab1cf-e381-442c-6e80-1ac2a8033b9f", options)

	// Boolean flag example
	fmt.Println("EnableTutorial's value is " + flags.EnableTutorial.IsEnabled(nil).)

	fmt.Println("This is experimental repo, have fun!")
	for {
	}
}

func forTest(a, b int) int {
	r := 10
	r = a + b*r
	if r == 10 {
		return r
	}
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	return r
}

type awscreds struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	IAMRole  string `json:"iam_role"`
	Password string `json:"password"`
}

func VStrings() {
	creds := &awscreds{
		Access:   "somekey",
		Secret:   "secret",
		Password: "myawesomepassword",
	}
	fmt.Println(creds)
}
