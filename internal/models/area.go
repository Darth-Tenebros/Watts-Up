package models

type Area struct {
	AreaName string `json:"area_name"`
	Finish   string `json:"fnish"`
	Source   string `json:"source"`
	Stage    int    `json:"stage"`
	Start    string `json:"start"`
}
