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
	// Update global window size
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	}

	// Get layout setting base on the terminal width and height
	layout := resolveLayout(m.termWidth, m.termHeight)
	
	// Handle all the update relate to dashboard
	nextDashboard, cmd, handledByDashboard := m.dashboard.Update(
		msg,
		len(m.filtered),
		layout.listWidth,
		layout.listHeight,
	)
	m.dashboard = nextDashboard
	m.refreshFilteredHosts()
	m.dashboard.SyncItemCount(len(m.filtered), layout.listHeight)

	// Handle Global hotkey for example exit the program
	if keyMsg, ok := msg.(tea.KeyPressMsg); ok && !handledByDashboard {
		key := keyMsg.String()
		switch {
		case slices.Contains(settings.Hotkey.Quit, key):
			return m, tea.Quit
		case slices.Contains(settings.Hotkey.Confirm, key):
			return m, m.enterSSHCard()
		}
	}

	return m, cmd
}

func (m Model) View() tea.View {
	layout := resolveLayout(m.termWidth, m.termHeight)
	selectedHost := m.selectedHost()

	// Render SSH list
	listPanel := m.dashboard.Render(m.filtered, layout.listWidth, layout.listHeight)

	// Render SSH detail sifebar
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

	// Render view and return
	v := tea.NewView(content)
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

func (m *Model) refreshFilteredHosts() {
	m.filtered = m.dashboard.FilterHosts(m.state.Hosts)
}
