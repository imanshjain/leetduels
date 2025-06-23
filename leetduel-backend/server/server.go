package server

import (
	"net/http"
	"leetduel-backend/router"
)

func New() *http.Server {
	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

// constant for firebase secret key
const SECRET string = "SECRETS/service-account-key.json"