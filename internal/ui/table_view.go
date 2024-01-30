package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type TableModel struct {
	Table table.Model
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "q", "ctrl+c":
			listModel := LoadListView()
			return listModel, nil
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m TableModel) View() string {
	return style.Render(m.Table.View()) + "\n"
}
