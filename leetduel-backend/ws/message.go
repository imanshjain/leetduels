package ws

import (
	"encoding/json"
	"log"

	"leetduel-backend/models"
)

type Message struct {
	Sender  *Client					`json:"sender"`
	Type    string          `json:"type"`
	Target  string          `json:"target"`
	Payload json.RawMessage	`json:"payload"`
}

type InboundMatchFound struct {
	P1 *Client `json:"p1"`
	P2 *Client `json:"p2"`
}

// Sent to the players after room is created
type OutboundMatchFound struct {
	RoomID     string `json:"roomId"`
	Question models.Question `json:"question"`
	TimeLimit  int    `json:"timeLimit"`
	Opponent   *string `json:"opponent"`
}

func mustMarshal(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic("Failed to marshal JSON: " + err.Error())
	}
	return b
}

func handleMessage(message Message, hub *Hub){
	switch message.Type{
	// Calls function to send the question and other details
	case "find_match":

		// TODO: Call the find players functions here.
		// TODO: Remove this
		dummyClient := &Client{
			conn: nil,                // no real WebSocket for now
			send: make(chan []byte), // just a dummy channel
			hub:  hub,    // pass the hub you're using
		}

		P1 := dummyClient
		P2 := dummyClient

		room := createNewRoom(P1, P2)

		hub.rooms[room] = true
		
		var outboundMessage OutboundMatchFound = OutboundMatchFound{
			RoomID: room.ID,
			Question: *room.question,
			TimeLimit: room.question.TimeLimit,
		}

		msg := Message{
			Type:    "match_found",
			Payload: mustMarshal(outboundMessage),
    }

    data, err := json.Marshal(msg)
    if err != nil {
			log.Println("Error marshaling match_found outbound message:", err)
			return
    }

    room.Player1.send <- data
    room.Player2.send <- data


	case "code_submit":
		// calls function to send the code to docker 
	case "disconnect":
		// closes the room and the websocket conneciton
	
	}
}