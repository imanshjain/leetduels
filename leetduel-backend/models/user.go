package models

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

// User struct represents a user in the system.
type User struct {
	Uid         string // Unique identifier for the user
	FirebaseUID string // Firebase UID for authentication
	Username    string // Username of the user
	Elo         int    // ELO rating of the user, default is 400
}

// GetUser returns a new User instance with default values.
func GetUser(ctx context.Context, token *auth.Token, dbpool *pgxpool.Pool) (*User, error) {

	// Get row from database
	row := dbpool.QueryRow(
		ctx,
		"SELECT uid, firebase_uid, username, elo FROM users WHERE firebase_uid = $1",
		token.UID,
	)

	var uid string
	var firebaseUID string
	var username string
	var elo int
	err := row.Scan(&uid, &firebaseUID, &username, &elo)

	if err != nil {
		return nil, fmt.Errorf("failed to get user from database: %w", err)
	}

	return &User{
		Uid:         uid,
		FirebaseUID: firebaseUID,
		Username:    username,
		Elo:         elo,
	}, nil
}
