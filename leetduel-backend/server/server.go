package server

import (
	"leetduel-backend/router"
	"net/http"

	"firebase.google.com/go/v4/auth"
)

func New(authClient *auth.Client) *http.Server {
	mux := http.NewServeMux()
	router.RegisterRoutes(mux, authClient)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
