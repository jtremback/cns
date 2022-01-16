package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GRecordKeyPrefix is the prefix to retrieve all GRecord
	GRecordKeyPrefix = "GRecord/value/"
)

// GRecordKey returns the store key to retrieve a GRecord from the index fields
func GRecordKey(
	key string,
) []byte {
	var k []byte

	keyBytes := []byte(key)
	k = append(k, keyBytes...)
	k = append(k, []byte("/")...)

	return k
}
