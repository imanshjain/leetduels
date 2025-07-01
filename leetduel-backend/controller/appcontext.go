package controller

import (
	"context"
	"leetduel-backend/utils"
	"leetduel-backend/ws"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

// AppContext holds the application context, which includes dependencies
type AppContext struct {
	AuthClient *auth.Client  // Firebase Auth client for user authentication
	Hub        *ws.Hub       // WebSocket hub for managing client connections
	Pgx        *pgxpool.Pool // PostgreSQL connection pool for database operations

	// Add any other application-wide dependencies here
}

func NewAppContext(authClient *auth.Client, hub *ws.Hub, pgxPool *pgxpool.Pool) *AppContext {
	// Initialize the AppContext with the required dependencies for middleware and handlers.
	return &AppContext{
		AuthClient: authClient,
		Hub:        hub,
		Pgx:        pgxPool,
	}
}

// Chain middlewares and handler
func (a_ctx *AppContext) ChainMiddleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// General purpose middleware go here. Tightly coupled middleware defined in their controller files.

// dbMiddleWare injects the database connection pool into the request context
// Add further proessing here to pool if needed.
func (a_ctx *AppContext) dbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Store the database connection pool in the request context
		newCtx := context.WithValue(r.Context(), utils.DbKey, a_ctx.Pgx)
		r = r.WithContext(newCtx)

		next(w, r)
	}
}

// hubMiddleware injects the WebSocket hub into the request context
// Add further processing here to hub if needed.
func (a_ctx *AppContext) hubMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Store the WebSocket hub in the request context
		newCtx := context.WithValue(r.Context(), utils.HubKey, a_ctx.Hub)
		r = r.WithContext(newCtx)

		next(w, r)
	}
}
