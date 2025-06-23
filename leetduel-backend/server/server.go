package server

import (
	"net/http"
	"leetduel-backend/router"
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

// constant for firebase secret key
const SECRET string = "SECRETS/service-account-key.json"