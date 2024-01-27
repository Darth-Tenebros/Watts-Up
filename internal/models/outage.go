package models

import "fmt"

type Outage struct {
	AreaName string `json:"area_name"`
	Finish   string `json:"finsh"`
	Source   string `json:"source"`
	Stage    int    `json:"stage"`
	Start    string `json:"start"`
}

func (o Outage) String() string {
	return fmt.Sprintf("AreaName: %s\nStart: %s\nFinish: %s\nStage: %d\n",
		o.AreaName, o.Start, o.Finish, o.Stage)
}
