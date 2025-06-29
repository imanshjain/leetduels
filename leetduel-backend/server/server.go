package server

import (
	"net/http"

	"firebase.google.com/go/v4/auth"

	"leetduel-backend/router"
	"leetduel-backend/ws"
)

func New(authClient *auth.Client) *http.Server {
	hub := ws.NewHub()
	go hub.Run()

	mux := http.NewServeMux()
	router.RegisterRoutes(mux, authClient, hub)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}