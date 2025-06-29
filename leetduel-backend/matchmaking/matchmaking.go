package matchmaking

// Needs client and room imports
import (
	"leetduel-backend/models"
)

// ELO bucket size const
const (
	ELOBucketSize = 25
)

// matchmaking queue that holds buckets of players
type MatchmakingQueue struct {
	Buckets map[int]*models.User // Map of ELO ratings to players
}

func NewMatchmakingQueue() *MatchmakingQueue {

	return &MatchmakingQueue{
		Buckets: make(map[int]*models.User), // Initialize the map of buckets
	}
}

func (mq *MatchmakingQueue) EnqueueOrGetRoom(player *models.User) *models.User {
	// Get bucket based on player's ELO rating
	bucketELO := player.Elo / ELOBucketSize
	bucket, exists := mq.Buckets[bucketELO]
	if !exists {
		// Create a new bucket if it doesn't exist
		mq.Buckets[bucketELO] = player

		return nil // No match found, enqueue player
	}

	// If a player exists in the bucket, return that player as a match
	delete(mq.Buckets, bucketELO) // Remove the matched player from the queue
	return bucket                 // Return the matched player
}

type MatchMakingInterface interface {
	EnqueueOrGetRoom(player *models.User, difficulty string) *models.User
}

type Matchmaking struct {
	EasyQueue   *MatchmakingQueue
	MediumQueue *MatchmakingQueue
	HardQueue   *MatchmakingQueue
}

func NewMatchmaking() *Matchmaking {
	return &Matchmaking{
		EasyQueue:   NewMatchmakingQueue(),
		MediumQueue: NewMatchmakingQueue(),
		HardQueue:   NewMatchmakingQueue(),
	}
}

func (m *Matchmaking) EnqueueOrGetRoom(player *models.User, difficulty string) *models.User {
	switch difficulty {
	case "easy":
		return m.EasyQueue.EnqueueOrGetRoom(player)
	case "medium":
		return m.MediumQueue.EnqueueOrGetRoom(player)
	case "hard":
		return m.HardQueue.EnqueueOrGetRoom(player)
	default:
		return nil // Invalid difficulty
	}
}
