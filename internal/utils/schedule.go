package utils

import (
	"encoding/json"
	"fmt"
	"watts-up/internal/models"
)

type Schedule struct {
	Times []models.Outage
}

func (sch *Schedule) UnmarshalResponse(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("no data to umarshal")
	}

	switch data[0] {
	case '{':
		return sch.unmarshalSingleItemResponse(data)
	case '[':
		return sch.unmarshalManyItemsResponse(data)
	}

	//jic
	err := sch.unmarshalManyItemsResponse(data)
	if err != nil {
		return sch.unmarshalSingleItemResponse(data)
	}
	return nil
}

func (sch *Schedule) unmarshalSingleItemResponse(data []byte) error {
	var outage models.Outage
	err := json.Unmarshal(data, &outage)
	if err != nil {
		return err
	}
	sch.Times = []models.Outage{outage}
	return nil
}

func (sch *Schedule) unmarshalManyItemsResponse(data []byte) error {
	var outages []models.Outage
	err := json.Unmarshal(data, &outages)
	if err != nil {
		return err
	}
	sch.Times = outages
	return nil
}
