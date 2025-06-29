package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Websocket connection for a specific client.

type Client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *Hub
}
const (
	pongWait = 60 * time.Second // How long to wait before a pong is received
	pingPeriod = (pongWait * 9) / 10 // Send a ping every pingPeriod seconds
	writeWait = 10 * time.Second // if more than this time is being used to write, the conneciton is considered to be broken
)

// goroutine that recieves information from the client
func (c *Client) readPump(){
	// Run at the end of the funciton/connection
	defer func() {
		// Unregister the client and close the conneciton
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))

	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	
	for {
		_, rawMsg, err := c.conn.ReadMessage()
    if err != nil {
			log.Println(err)
			return
    }
		// The message is sent to hub who figures out what to do with it.
    // var msg Message
		// if err := json.Unmarshal(messageBytes, &msg); err != nil {
		// 	log.Println("Invalid message:", err)
		// 	continue
		// }
	
		var msg Message
		if err := json.Unmarshal(rawMsg, &msg); err != nil {
			log.Println("invalid message format:", err)
			continue
		}

		// Optionally attach the client info to the message if needed
		msg.Sender = c // if you have Sender *Client field in Message

		c.hub.broadcast <- msg
	}
}

// goroutine that sends information to the client
func (c *Client) writePump(){
	// Send a ping message to the client every pingPeriod
	ticker := time.NewTicker(pingPeriod)
	defer func ()  {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
    select {
			// Read the message from the send channel.
			// ok = true if the channel is open and you got a message.
			case msg, ok := <-c.send:
				if !ok {
					// The hub has closed the channel
					c.conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				c.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					return
				}

			// Checks the ticker channel to see when it is time to send ping
			case <-ticker.C:
				c.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
    }
	}
}
