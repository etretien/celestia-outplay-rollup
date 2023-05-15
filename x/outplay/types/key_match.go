package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MatchKeyPrefix is the prefix to retrieve all Match
	MatchKeyPrefix = "Match/value/"
)

// MatchKey returns the store key to retrieve a Match from the index fields
func MatchKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
