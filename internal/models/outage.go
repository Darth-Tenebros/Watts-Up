package models

import (
	"fmt"
)

// Outage represents /(maps to) a load shedding outage as defined by the API
type Outage struct {
	AreaName string `json:"area_name"`
	Finish   string `json:"finsh"`
	Source   string `json:"source"`
	Stage    int    `json:"stage"`
	Start    string `json:"start"`
}

func (out Outage) String() string {
	return fmt.Sprintf("AreaName: %s\nStart: %s\nFinish: %s\nStage: %d\n",
		out.AreaName, out.Start, out.Finish, out.Stage)
}
