/*
Package utils has utilities for the project
*/
package utils

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
)

// GenerateId generates a unique sha256 id with the given area name
func GenerateId(areaName string) string {
	hash := sha256.Sum256([]byte(areaName))
	return hex.EncodeToString(hash[:])
}

// OpenDatabase returns a pointer for an sql.DB object
// the object connects to the SQLite3 DB using the mattn SQLite3 driver
func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./store/favourites")
	if err != nil {
		return nil, err
	}
	return db, nil
}
