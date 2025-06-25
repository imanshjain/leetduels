package main

import (
	"fmt"
	"log"

	"leetduel-backend/db"
	"leetduel-backend/server"
)

func main() {

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
		fmt.Println(err)
		log.Fatal("❌ Error initializing firebase")
	}
	fmt.Println("Firebase Auth initialized.")

	// Create and run new server.
	s := server.New(authClient)
	log.Fatal(s.ListenAndServe())
}
