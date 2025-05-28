package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/movie-rental-app/util"
)

func TestENVRead(t *testing.T) {
	fmt.Println("testing works")

	err := godotenv.Load("../config-local.env")
	if err != nil {
		t.Errorf("Error loading .env file: %v", err)
	}

	// Assersion check to check the hostname
	util.Assert(t, "host.docker.internal", os.Getenv("DB_HOST"))
	fmt.Println("Environment variables loaded successfully")
}
