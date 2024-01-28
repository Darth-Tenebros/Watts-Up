package main

import (
	"eskom-se-poes/internal/ui"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	link := "https://eskom-calendar-api.shuttleapp.rs/list_areas"
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	areas := strings.Split(strings.TrimSuffix(string(body), "\n"), ",")
	items := ui.AreasToListItems(areas)

	m := ui.ListModel{List: list.New(items, list.NewDefaultDelegate(), 10, 10)}
	m.List.Title = "SA Load Shedding Areas"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
