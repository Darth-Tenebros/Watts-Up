/*
Package repository implements a library for communicating with the DB
*/
package repository

import (
	"database/sql"
	"fmt"
	"watts-up/internal/utils"
)

// AreaRepo is implements the driver for connecting to the DB
// uses the std lib database/sql
type AreaRepo struct {
	DB *sql.DB
}

// NewAreaRepo returns a pointer to a new instance of AreaRepo
func NewAreaRepo(db *sql.DB) *AreaRepo {
	return &AreaRepo{
		DB: db,
	}
}

// AddFavourite adds the given area (name) to the DB
// returns the unique id of the added area or an error
func (area *AreaRepo) AddFavourite(areaName string) (string, error) {
	id := utils.GenerateId(areaName)

	exists, err := area.idExists(id)
	if err != nil {
		return "", err
	}

	if !exists {
		sqlStatement, err := area.DB.Prepare("INSERT INTO areas(id, area_name) VALUES (?, ?)")
		if err != nil {
			return "", err
		}
		defer sqlStatement.Close()

		_, err = sqlStatement.Exec(id, areaName)
		if err != nil {
			return "", err
		}
		return id, nil
	}
	return "", fmt.Errorf("the area is already in the favourites database")
}

// idExists checks whether the given string id exists in the DB
// returns bool or error
func (area *AreaRepo) idExists(id string) (bool, error) {
	var count int
	err := area.DB.QueryRow("SELECT COUNT(*) FROM areas WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetAllAreaNames retrieves all the area names from the DB
// returns a slice with the area names or an error
func (area *AreaRepo) GetAllAreaNames() ([]string, error) {
	var names []string
	rows, err := area.DB.Query("SELECT area_name FROM areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var areaName string
		if err := rows.Scan(&areaName); err != nil {
			return nil, err
		}
		names = append(names, areaName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return names, nil
}

// DeleteAreaFromFavourites deletes the area with the given id from the DB
// returns an error if something goes wrong
func (area *AreaRepo) DeleteAreaFromFavourites(id string) error {
	sqlStatement, err := area.DB.Prepare("DELETE FROM areas WHERE id = ?")
	if err != nil {
		return err
	}
	defer sqlStatement.Close()

	_, err = sqlStatement.Exec(id)
	return err
}
