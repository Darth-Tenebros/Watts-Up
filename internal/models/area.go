package models

import "fmt"

// Area represents an area that can be affected by load shedding
// This entity helps with DB interaction
type Area struct {
	Id       string
	AreaName string
}

func (a Area) String() string {
	return fmt.Sprintf("Area: %s", a.AreaName)
}
