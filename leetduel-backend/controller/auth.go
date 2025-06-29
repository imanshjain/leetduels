package controller

import (
	"context"
	"fmt"
	"net/http"

	"firebase.google.com/go/v4/auth"
)

func verifyIDToken(authClient *auth.Client, ctx context.Context, idToken string) (*auth.Token, error) {
	// Verify the ID token
	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil

}

func GetUser(w http.ResponseWriter, r *http.Request, authClient *auth.Client) {
	ctx := context.Background()

	fmt.Println("GetUser called")

	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	token, err := verifyIDToken(authClient, ctx, idToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID token: %v", err), http.StatusUnauthorized)
		return
	}

	// Write the user profile information to the response
	fmt.Fprintf(w, "User ID: %s\nEmail: %s\n", token.UID, token.Claims["email"])
}
