package models

import (
	"fmt"
	"strconv"
)

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

func (out Outage) outageToSlice(outage Outage) []string {
	return []string{strconv.Itoa(out.Stage), out.AreaName, out.Start, out.Finish}
}
