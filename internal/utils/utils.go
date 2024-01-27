package utils

import (
	"encoding/json"
	"eskom-se-poes/internal/models"
)

type Schedule struct {
	times []models.Outage
}

func (sch *Schedule) unmarshalSingleItemResponse(data []byte) error {
	var outage models.Outage
	err := json.Unmarshal(data, &outage)
	if err != nil {
		return err
	}
	sch.times = []models.Outage{outage}
	return nil
}

func (sch *Schedule) unmarshalManyItemsResponse(data []byte) error {
	var outages []models.Outage
	err := json.Unmarshal(data, &outages)
	if err != nil {
		return err
	}
	sch.times = outages
	return nil
}
