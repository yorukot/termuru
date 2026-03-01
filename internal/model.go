package internal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	count int
}

func IntinalModel() Model {
	return Model{count: 0}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		case "down":
			m.count--
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	s := fmt.Sprintf("Count: %d\nPress up/down to change. q to quit.\n", m.count)

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}
