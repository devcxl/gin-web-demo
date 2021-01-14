package common

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
获取message的sha256
*/
func GetSha256Code(message []byte) string {
	sum256 := sha256.Sum256(message)
	hashcode := hex.EncodeToString(sum256[:])
	return hashcode
}
