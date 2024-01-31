package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateId(areaName string) string {
	hash := sha256.Sum256([]byte(areaName))
	return hex.EncodeToString(hash[:])
}
