package main

import (
	"eskom-se-poes/internal/ui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	listModel := ui.LoadListView()
	p := tea.NewProgram(listModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
