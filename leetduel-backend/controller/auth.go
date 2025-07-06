package controller

import (
	"context"
	"fmt"
	"leetduel-backend/utils"
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

func (a_ctx *AppContext) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		fmt.Println("AuthMiddleware called")

		// Use Auth header with HTTP, Query with WebSocket. Use WSS for encrypting the ID token.
		idToken := r.Header.Get("Authorization")
		if idToken == "" {
			idToken = r.URL.Query().Get("idToken")

			if idToken == "" {
				http.Error(w, "Unauthorized: ID token is missing", http.StatusUnauthorized)
				return
			}
		}

		token, err := verifyIDToken(a_ctx.AuthClient, ctx, idToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid ID token: %v", err), http.StatusUnauthorized)
			return
		}

		// Store the user ID in the request context for further processing
		newCtx := context.WithValue(r.Context(), utils.AuthKey, token)
		r = r.WithContext(newCtx)

		next(w, r)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	// Retrieve the token from the request context
	token, ok := r.Context().Value(utils.AuthKey).(*auth.Token)

	if !ok || token == nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Write the user profile information to the response
	fmt.Fprintf(w, "User ID: %s\nEmail: %s\n", token.UID, token.Claims["email"])
}
