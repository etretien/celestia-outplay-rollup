package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChallengeKeyPrefix is the prefix to retrieve all Challenge
	ChallengeKeyPrefix = "Challenge/value/"
)

// ChallengeKey returns the store key to retrieve a Challenge from the index fields
func ChallengeKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
