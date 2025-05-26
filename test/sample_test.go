package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestENVRead(t *testing.T) {
	fmt.Println("testing works")

	err := godotenv.Load("../config-local.env")
	if err != nil {
		t.Errorf("Error loading .env file: %v", err)
	}

	// Assersion check to check the hostname
	assert(t, "host.docker.internal", os.Getenv("DB_HOST"))
	fmt.Println("Environment variables loaded successfully")
}

func assert(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	} else {
		fmt.Println("Assertion passed")
	}
}
