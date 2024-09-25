package view

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	docStyle          = lipgloss.NewStyle().Margin(1, 2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

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

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Choice)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Title())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list   list.Model
	choice interface{}
}

func CreateView(question string, choiceResult []Choice) model {
	log.Printf("%v", choiceResult[0].Title())

	items := []list.Item{
		choiceResult[0],
		choiceResult[1],
		choiceResult[2],
		choiceResult[3],
	}

	newList := list.New(items, itemDelegate{}, 0, 0)
	newList.Title = question
	newList.SetShowStatusBar(false)
	newList.SetFilteringEnabled(false)
	newList.Styles.Title = titleStyle

	return model{
		list:   newList,
		choice: nil,
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
		case "enter":
			i, ok := m.list.SelectedItem().(Choice)
			if ok {
				m.choice = i
			}
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
	if option, ok := m.choice.(Choice); ok {
		if option.Correct {
			return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", option.Title()))
		}
		return quitTextStyle.Render(fmt.Sprintf("It's not %s.", option.Title()))
	} else {
		log.Println("m.Choice is not an Choice struct.")
	}
	return docStyle.Render(m.list.View())
}
