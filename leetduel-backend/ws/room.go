package ws

import (
	"time"

	"github.com/google/uuid"

	"leetduel-backend/models"
)

type Room struct {
	ID string

	Player1 *Client

	Player2 *Client

	question *models.Question

	StartTime time.Time

	Timelimit *time.Time
}

func generateRoomID() string {
	return uuid.New().String()
}

func createNewRoom(p1 *Client, p2 *Client) *Room {
	var que *models.Question = models.GetQuestion()
	now := time.Now()
	timelimit := now.Add(time.Duration(que.TimeLimit) * time.Minute)

	return &Room{
		ID:        generateRoomID(),
		Player1:   p1,
		Player2:   p2,
		question:  que,
		StartTime: now,
		Timelimit: &timelimit,
	}
}
