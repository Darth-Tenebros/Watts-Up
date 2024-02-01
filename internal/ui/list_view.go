package ui

import (
	"eskom-se-poes/internal/repository"
	"eskom-se-poes/internal/utils"
	"log"
	"regexp"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	AreaName string
}

func AreasToListItems(areas []string) []list.Item {
	var items []list.Item
	for _, area := range areas {
		item := Item{
			AreaName: area,
		}
		items = append(items, item)
	}
	return items
}

func (i Item) Title() string       { return i.AreaName }
func (i Item) Description() string { return "random desc" }
func (i Item) FilterValue() string { return i.AreaName }

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ListModel struct {
	List         list.Model
	SelectedItem int
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			areaName := m.List.Items()[m.List.Index()]
			re := regexp.MustCompile("[^a-zA-Z0-9-]")
			return TableModel{Table: LoadTableView(re.ReplaceAllString(areaName.FilterValue(), ""))}, nil

		case "a":
			areaName := m.List.Items()[m.List.Index()]
			re := regexp.MustCompile("[^a-zA-Z0-9-]")

			db, err := utils.OpenDatabase()
			if err != nil {
				return m, nil
			}
			defer db.Close()

			areaRepo := repository.NewAreaRepo(db)
			_, err = areaRepo.AddFavourite(re.ReplaceAllString(areaName.FilterValue(), ""))
			log.Println(err)
			if err != nil {
				return m, nil
			}
			return m, nil

		case "b":
			listModel := LoadListView()
			return listModel, nil

		case "f":
			db, err := utils.OpenDatabase()
			if err != nil {
				return m, nil
			}
			defer db.Close()
			areaRepo := repository.NewAreaRepo(db)

			names, err := areaRepo.GetAllAreaNames()
			if err != nil {
				log.Println(err)
				return m, nil
			}
			items := AreasToListItems(names)
			m := ListModel{List: list.New(items, list.NewDefaultDelegate(), 10, 10)}
			//m.List.SetHeight(35)
			//m.List.SetWidth(50)
			m.List.Title = "SA Load Shedding Areas - FAVOURITES"
			return m, nil

		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
func (m ListModel) View() string {
	m.List.SetHeight(35)
	m.List.SetWidth(50)
	return docStyle.Render(m.List.View())
}
