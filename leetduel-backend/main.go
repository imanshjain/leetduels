package main

import (
	"fmt"
	"log"

	"github.com/lpernett/godotenv"

	"leetduel-backend/db"
	"leetduel-backend/server"
)

func main() {
	// Checking if the dot file loads correctly
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Initiliaze pgql connection
  dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal("❌ Error connecting to pg")
	}
	defer dbPool.Close()
	fmt.Println("PostgreSQL initialized.")

	// Initilaize Firebase connection.
	authClient, err := db.InitializeFirebase()
	if err != nil {
		log.Fatal("❌ Error initializing firebase")
	}
	fmt.Println("Firebase Auth initialized.")

	// Create and run new server.
	s := server.New(authClient)
	log.Fatal(s.ListenAndServe())
}
