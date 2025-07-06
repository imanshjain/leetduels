package controller

import (
	"context"
	"leetduel-backend/models"
	"leetduel-backend/utils"
	"leetduel-backend/ws"
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitMessage(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	token, ok_t := r.Context().Value(utils.AuthKey).(*auth.Token)
	db, ok_d := r.Context().Value(utils.DbKey).(*pgxpool.Pool)

	log.Printf("InitMessage called with token: %v, db: %v", ok_t, ok_d)

	if !ok_t || !ok_d || token == nil || db == nil {
		http.Error(w, "Unauthorized: Invalid token or context", http.StatusUnauthorized)
		return
	}

	user, err := models.GetUser(ctx, token, db)

	if err != nil {
		http.Error(w, "Failed to get user for token", http.StatusInternalServerError)
		return
	}

	newCtx := context.WithValue(r.Context(), utils.UserKey, user)
	r = r.WithContext(newCtx)

	// Promote to ws
	err_ := ws.ServeWS(w, r)
	if err_ != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
}
