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

func (mq *MatchmakingQueue) EnqueueOrGetRoom(player *models.User) *models.User {
	return nil
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
