package server

import (
	"leetduel-backend/router"
	"net/http"

	"leetduel-backend/controller"
)

func New(a_ctx *controller.AppContext) *http.Server {
	mux := http.NewServeMux()
	router.RegisterRoutes(a_ctx, mux)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
