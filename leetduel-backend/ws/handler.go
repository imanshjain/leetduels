package ws

import (
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

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// Step 1: Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}

	// Step 2: Create a new client
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256), // Buffered channel to avoid blocking
	}

	// Step 3: Register the client with the hub
	client.hub.register <- client

	// Step 4: Start read and write goroutines
	go client.writePump()
	go client.readPump()
}