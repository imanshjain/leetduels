package db

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/option"
)

func ConnectDB() (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("DB URL:", dbURL)

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return dbpool, nil
}

// Initialize Firebase Admin SDK
func InitializeFirebase() (*auth.Client, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SECRET_KEY"))
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	return app.Auth(ctx)
}
