package models

import "fmt"

type Area struct {
	Id       string
	AreaName string
}

func (a Area) String() string {
	return fmt.Sprintf("Area: %s", a.AreaName)
}
