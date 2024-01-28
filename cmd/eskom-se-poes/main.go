package main

import (
	"eskom-se-poes/internal/ui"
	"eskom-se-poes/internal/utils"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"net/http"
	"os"
)

func main() {

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

func loadTableView() {
	link := "https://eskom-calendar-api.shuttleapp.rs/outages/"
	// TODO: FETCH FROM COMMAND LINE
	location := "city-of-cape-town-area-15"

	schedule, err := getSchedule(link, location)
	if err != nil {
		fmt.Print(err)
	}

	var rows []table.Row
	for _, outage := range schedule.Times {
		rows = append(rows, outage.OutageToSlice())
	}

	columns := []table.Column{
		{Title: "Stage", Width: 5},
		{Title: "Area Name", Width: 30},
		{Title: "Start", Width: 30},
		{Title: "Finish", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := ui.TableModel{Table: t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
