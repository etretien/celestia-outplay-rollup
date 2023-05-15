package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProfileKeyPrefix is the prefix to retrieve all Profile
	ProfileKeyPrefix = "Profile/value/"
)

// ProfileKey returns the store key to retrieve a Profile from the index fields
func ProfileKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
