package view

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"jobnbackpack.com/answer_generator/chat"
	"jobnbackpack.com/answer_generator/models"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	list        list.Model
	choice      interface{}
	input       textinput.Model
	currentView int
}

func CreateView() model {
	ti := textinput.New()
	ti.Placeholder = "Welcher der Juenger konnte auf Wasser gehen?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		currentView: 0,
		input:       ti,
		choice:      nil,
	}
}

func (m model) Init() tea.Cmd {
	switch m.currentView {
	case 0:
		return textinput.Blink
	case 1:
		return nil
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.currentView {
			case 1:
				i, ok := m.list.SelectedItem().(models.Choice)
				if ok {
					m.choice = i
				}
				// return m, tea.Quit
			case 0:
				question := m.input.Value()
				newList := CreateList(question, chat.AskGPT(question))
				m.list = newList
				m.currentView = 1
			}
		}
	case tea.WindowSizeMsg:
		if m.currentView == 1 {
			h, v := docStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v)
		}

	}

	var cmd tea.Cmd
	switch m.currentView {
	case 1:
		m.list, cmd = m.list.Update(msg)
	case 0:
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	switch m.currentView {
	case 1:
		if option, ok := m.choice.(models.Choice); ok {
			if option.Correct {
				return quitTextStyle.Render(fmt.Sprintf("%s ist richtig!", option.Title()))
			}
			return quitTextStyle.Render(fmt.Sprintf("Es ist nicht %s.", option.Title()))
		} else {
			log.Println("Choice is not an Choice struct.")
		}
		return docStyle.Render(m.list.View())
	case 0:
		return fmt.Sprintf(
			"Stelle eine Quiz Frage zur Bibel:\n\n%s\n\n%s",
			m.input.View(),
			"(esc to quit)",
		) + "\n"
	}
	return quitTextStyle.Render("Es gibt keine current view")
}
