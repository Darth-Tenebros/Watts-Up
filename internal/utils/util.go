package utils

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
)

func GenerateId(areaName string) string {
	hash := sha256.Sum256([]byte(areaName))
	return hex.EncodeToString(hash[:])
}

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./store/favourites")
	if err != nil {
		return nil, err
	}
	return db, nil
}
