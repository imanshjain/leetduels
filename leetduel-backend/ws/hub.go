package ws

// This file contains code to keep track of everything.

// Responsibilities:
// 	•	Keeps track of all connected clients
// 	•	Handles registration/unregistration of clients
// 	•	Routes messages to the correct recipients
// 	•	(Eventually) does matchmaking or coordinates with Redis

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}