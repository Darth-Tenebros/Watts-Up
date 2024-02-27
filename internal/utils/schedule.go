/*
Package utils has the utility functions of the project
*/
package utils

import (
	"encoding/json"
	"fmt"
	"watts-up/internal/models"
)

// Schedule has a slice of outages (for some area)
type Schedule struct {
	Times []models.Outage
}

// UnmarshalResponse maps the json response to an Outage object
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

// unmarshalSingleItemResponse maps a single json response to an Outage into a Schedule
// slice of size 1 with the outage
func (sch *Schedule) unmarshalSingleItemResponse(data []byte) error {
	var outage models.Outage
	err := json.Unmarshal(data, &outage)
	if err != nil {
		return err
	}
	sch.Times = []models.Outage{outage}
	return nil
}

// unmarshalManyItemsResponse maps a list (>= 2) of json items to a slice of Outages in the Schedule object
func (sch *Schedule) unmarshalManyItemsResponse(data []byte) error {
	var outages []models.Outage
	err := json.Unmarshal(data, &outages)
	if err != nil {
		return err
	}
	sch.Times = outages
	return nil
}
