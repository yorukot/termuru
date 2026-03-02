package internal

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/yorukot/termuru/internal/settings"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		key := msg.String()
		switch {
		case slices.Contains(settings.Hotkey.Quit, key):
			return m, tea.Quit
		case slices.Contains(settings.Hotkey.Select, key):
			return m, m.enterSSHCard()
		}
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	}

	layout := resolveLayout(m.termWidth, m.termHeight)
	nextDashboard, cmd := m.dashboard.Update(
		msg,
		len(m.state.Hosts),
		layout.listWidth,
		layout.listHeight,
	)
	m.dashboard = nextDashboard

	return m, cmd
}

func (m Model) View() tea.View {
	layout := resolveLayout(m.termWidth, m.termHeight)
	selectedHost := m.selectedHost()

	listPanel := m.dashboard.Render(m.state.Hosts, layout.listWidth, layout.listHeight)

	var content string
	switch layout.placement {
	case placementRight:
		infoPanel := m.sshInfo.Render(selectedHost, layout.infoWidth, layout.infoHeight)
		content = lipgloss.JoinHorizontal(lipgloss.Top, listPanel, infoPanel)
	case placementBottom:
		infoPanel := m.sshInfo.Render(selectedHost, layout.infoWidth, layout.infoHeight)
		content = lipgloss.JoinVertical(lipgloss.Left, listPanel, infoPanel)
	default:
		content = listPanel
	}

	v := tea.NewView(content)
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
