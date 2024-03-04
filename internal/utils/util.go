/*
Package utils has utilities for the project
*/
package utils

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

// GenerateId generates a unique sha256 id with the given area name
func GenerateId(areaName string) string {
	hash := sha256.Sum256([]byte(areaName))
	return hex.EncodeToString(hash[:])
}

// OpenDatabase returns a pointer for an sql.DB object
// the object connects to the SQLite3 DB using the mattn SQLite3 driver
func OpenDatabase() (*sql.DB, error) {
	// Find the project directory by traversing upwards until we find a known marker file or directory
	projectDir, err := findProjectDir()
	if err != nil {
		return nil, err
	}

	// Construct the absolute path to the database file
	dbPath := filepath.Join(projectDir, "store/favourites") // Adjust the relative path as needed

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// some hacky bs to traverse dirs because form some reason, the db dir worked fine in the goland ide but not
// in the system command line
func findProjectDir() (string, error) {
	// Start from the current directory of the file containing this function
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse upwards until we find a known marker file or directory
	for {
		if isProjectDir(dir) {
			return dir, nil
		}

		// Move up one directory
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory
			return "", errors.New("project directory not found")
		}
		dir = parentDir
	}
}

func isProjectDir(dir string) bool {
	// Check if the directory contains a known marker file or directory
	_, err := os.Stat(filepath.Join(dir, "go.mod")) // Check for the presence of go.mod file
	return err == nil
}
