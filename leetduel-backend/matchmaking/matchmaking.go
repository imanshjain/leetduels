package matchmaking

// Needs client and room imports
import (
	"leetduel-backend/models"
)

// ELO bucket size const
const (
	ELOBucketSize = 25
)

// Bucket represents a matchmaking bucket for players with similar ELO ratings.
type Bucket struct {
	Players []*models.User // List of players in the bucket
}

// matchmaking queue that holds buckets of players
type MatchmakingQueue struct {
	Buckets map[int]*Bucket // Map of ELO ratings to buckets
}

func NewMatchmakingQueue() *MatchmakingQueue {

	return &MatchmakingQueue{
		Buckets: make(map[int]*Bucket),
	}
}
