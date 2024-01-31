package main

import (
	"eskom-se-poes/internal/ui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func initLogger() {
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)
}

func main() {

	initLogger()

	listModel := ui.LoadListView()
	p := tea.NewProgram(listModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
