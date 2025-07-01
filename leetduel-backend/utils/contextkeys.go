package utils

type contextKey string

const (
	AuthKey contextKey = "authKey" // authMiddleware context key
	DbKey   contextKey = "dbKey"   // dbMiddleware context key
	HubKey  contextKey = "hubKey"  // hubMiddleware context key
	UserKey contextKey = "userKey" // user context key
)
