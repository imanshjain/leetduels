package main

import (
	"fmt"
	"log"

	"leetduel-backend/controller"
	"leetduel-backend/db"
	"leetduel-backend/server"
	"leetduel-backend/ws"
)

func main() {

	// Initialize pgql connection
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

	// Initialize WebSocket hub.
	hub := ws.NewHub()
	go hub.Run()

	// Initialize AppContext with auth client and hub.
	a_ctx := controller.NewAppContext(authClient, hub, dbPool)

	// Create and run new server.
	s := server.New(a_ctx)
	log.Fatal(s.ListenAndServe())
}
