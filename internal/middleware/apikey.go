package middleware

import (
	"os"
	"strings"
)

// AccessTier defines the rate limit tier for a request.
type AccessTier int

const (
	// TierAnonymous is used for requests without a valid API key.
	TierAnonymous AccessTier = iota
	// TierAPIKey is used for requests with a valid API key.
	TierAPIKey
)

// APIKeyService holds the set of valid API keys loaded from environment.
// Designed for future migration to a database-backed store.
type APIKeyService struct {
	validKeys map[string]struct{}
}

// NewAPIKeyService creates an APIKeyService by reading the API_KEYS
// environment variable (comma-separated list of keys).
// Example: API_KEYS=internal123,partner456,test789
func NewAPIKeyService() *APIKeyService {
	keys := make(map[string]struct{})

	rawKeys := os.Getenv("API_KEYS")
	if rawKeys != "" {
		for _, k := range strings.Split(rawKeys, ",") {
			trimmed := strings.TrimSpace(k)
			if trimmed != "" {
				keys[trimmed] = struct{}{}
			}
		}
	}

	return &APIKeyService{validKeys: keys}
}

// IsValid reports whether the given API key exists in the key set.
func (s *APIKeyService) IsValid(key string) bool {
	_, ok := s.validKeys[key]
	return ok
}

// HasKeys reports whether any API keys have been configured.
func (s *APIKeyService) HasKeys() bool {
	return len(s.validKeys) > 0
}
