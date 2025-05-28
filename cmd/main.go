package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/movie-rental-app/app"
)

func main() {
	fmt.Println("Starting Movie Rental App...")
	//create a APP
	err := godotenv.Load("config-local.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	app := app.NewApp()

	// initialize the app
	err = app.Init()
	if err != nil {
		panic(fmt.Sprintf("Error initializing app: %v", err))
	}
	fmt.Println("App initialized successfully")

}
