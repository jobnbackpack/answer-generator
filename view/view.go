package view

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Choice struct {
	Value   string `json:"value"`
	Correct bool   `json:"correct"`
}

func (c Choice) FilterValue() string {
	return c.Value
}

func (c Choice) Title() string {
	return c.Value
}

type model struct {
	list list.Model
}

func CreateView(question string, choiceResult []Choice) model {
	items := []list.Item{
		choiceResult[0],
		choiceResult[1],
		choiceResult[2],
	}

	newList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	newList.Title = question

	return model{
		list: newList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
