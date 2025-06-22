import (
	"context"
	"fmt"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"leetduel-backend/server"
)

// Initialize Firebase Admin SDK
func initializeFirebase() (*auth.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("../SECRETS/service-account-key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	return app.Auth(ctx)
}

// Server is the main entry point for the server application.
func main() {
	s := server.New()
  log.Fatal(s.ListenAndServe())

	_, err := initializeFirebase()
	if err != nil {
		fmt.Println("Error initializing LeetDuel server:", err)
		return
	}

	fmt.Println("Firebase Auth initialized.")
}
