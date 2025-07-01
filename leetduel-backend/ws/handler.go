package ws

import (
	"leetduel-backend/models"
	"leetduel-backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow connections from any origin for dev purposes.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(w http.ResponseWriter, r *http.Request) error {
	// Step 1: Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return err
	}

	// Step 2: Create a new client
	hub, ok_h := r.Context().Value(utils.HubKey).(*Hub)
	user, ok_u := r.Context().Value(utils.UserKey).(*models.User)

	if !ok_h || !ok_u || hub == nil || user == nil {
		log.Println("Hub or user context is missing")
	}

	client := NewClient(conn, hub, user)

	// Step 3: Register the client with the hub
	client.hub.register <- client

	// Step 4: Start read and write goroutines
	go client.writePump()
	go client.readPump()

	return nil
}
