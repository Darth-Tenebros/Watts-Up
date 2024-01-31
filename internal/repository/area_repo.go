package repository

import (
	"database/sql"
	"eskom-se-poes/internal/utils"
	"fmt"
)

type AreaRepo struct {
	DB *sql.DB
}

func NewAreaRepo(db *sql.DB) *AreaRepo {
	return &AreaRepo{
		DB: db,
	}
}
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

func (area *AreaRepo) idExists(id string) (bool, error) {
	var count int
	err := area.DB.QueryRow("SELECT COUNT(*) FROM areas WHERE id = ?", id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
