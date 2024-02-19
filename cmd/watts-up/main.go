package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"watts-up/internal/ui"
)

func main() {
	listModel := ui.LoadListView()
	p := tea.NewProgram(listModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
