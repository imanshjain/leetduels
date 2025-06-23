package main

import (
	"context"
	"fmt"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
	"leetduel-backend/server"
)

// Initialize Firebase Admin SDK
func initializeFirebase() (*auth.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(server.SECRET)
	app, err := firebase.NewApp(ctx, nil, opt)
	
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	return app.Auth(ctx)
}

// Server is the main entry point for the server application.
func main() {

	authClient, err := initializeFirebase()
	if err != nil {
		fmt.Println("Error initializing LeetDuel server:", err)
		return
	}

	s := server.New(authClient)

	fmt.Println("Firebase Auth initialized.")

	log.Fatal(s.ListenAndServe())
}
