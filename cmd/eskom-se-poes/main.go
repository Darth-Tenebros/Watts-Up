package main

import (
	"eskom-se-poes/internal/utils"
	"fmt"
	"io"
	"net/http"
)

func main() {
	link := "https://eskom-calendar-api.shuttleapp.rs/outages/"
	location := "city-of-cape-town-area-15"

	schedule, err := getSchedule(link, location)
	if err != nil {
		fmt.Print(err)
	}

	for _, outage := range schedule.Times {
		fmt.Println(outage)
	}
}

func getSchedule(link, area string) (*utils.Schedule, error) {
	fullUrl := link + area
	var schedule utils.Schedule

	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = schedule.UnmarshalResponse(body)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}
