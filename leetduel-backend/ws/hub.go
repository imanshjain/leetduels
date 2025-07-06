package ws

import (
	"fmt"
	"leetduel-backend/matchmaking"
	"log"
)

// This file contains code to keep track of everything.

// Responsibilities:
// 	•	Keeps track of all connected clients
// 	•	Handles registration/unregistration of clients
// 	•	Routes messages to the correct recipients
// 	•	(Eventually) does matchmaking or coordinates with Redis

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Active Rooms.
	rooms map[*Room]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Map of uid to client
	uidToClient map[string]*Client

	// Matchmaking interface
	matchmaking matchmaking.MatchMakingInterface
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan Message),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		rooms:       make(map[*Room]bool),
		uidToClient: make(map[string]*Client),
	}
}

func (hub *Hub) Run() {

	fmt.Println("WebSocket hub is running...")

	for {
		select {
		case client := <-hub.register:
			log.Printf("Client registered: %s\n", client.User.Uid)
			hub.clients[client] = true
			hub.uidToClient[client.User.Uid] = client // Map the client's UID to the client

		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				delete(hub.uidToClient, client.User.Uid) // Remove the client from the UID map
				close(client.send)
			}

		case message := <-hub.broadcast:
			// Send the message/code to the docker container to run the code with the tests.
			fmt.Printf("Code to be run was received: %s\n", message)
			handleMessage(message, hub)

		}
	}
}
