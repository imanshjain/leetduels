package ws

import "fmt"

// This file contains code to keep track of everything.

// Responsibilities:
// 	•	Keeps track of all connected clients
// 	•	Handles registration/unregistration of clients
// 	•	Routes messages to the correct recipients
// 	•	(Eventually) does matchmaking or coordinates with Redis

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (hub *Hub) run() {
	for {
		select {
			case client := <-hub.register:
				hub.clients[client] = true

			case client := <-hub.unregister:
				if _, ok := hub.clients[client]; ok {
					delete(hub.clients, client)
					close(client.send)
				}

			case message := <-hub.broadcast:
				// Send the message/code to the docker container to run the code with the tests.
				fmt.Printf("Code to be run was received: %s\n", message)
		}
	}
}

