package server

import (
	"net/http"
	"leetduel-backend/router/router"
)

func New() *http.Server {
	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}