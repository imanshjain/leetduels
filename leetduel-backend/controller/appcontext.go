package controller

import (
	"leetduel-backend/ws"
	"net/http"

	"firebase.google.com/go/v4/auth"
)

// AppContext holds the application context, including the Firebase Auth client and WebSocket hub.
type AppContext struct {
	AuthClient *auth.Client // Firebase Auth client for user authentication
	Hub        *ws.Hub      // WebSocket hub for managing client connections

	// Add any other application-wide dependencies here
}

type contextKey string

const (
	authKey contextKey = "authKey" // authMiddleware context key
)

func NewAppContext(authClient *auth.Client, hub *ws.Hub) *AppContext {
	// Initialize the AppContext with the provided Firebase Auth client and WebSocket hub
	return &AppContext{
		AuthClient: authClient,
		Hub:        hub,
	}
}

// Chain middlewares and handler
func (a_ctx *AppContext) ChainMiddleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
