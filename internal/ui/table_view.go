package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("240"))

// TableModel represents a list.Item entry
type TableModel struct {
	Table table.Model
}

// Init implements tea.Model interface for TableModel
func (m TableModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model interface for TableModel
// handles the click / interaction events of the tui
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
		case "b":
			listModel := LoadListView()
			return listModel, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

// View implements tea.Model interface for TableModel
func (m TableModel) View() string {
	return style.Render(m.Table.View()) + "\n"
}
