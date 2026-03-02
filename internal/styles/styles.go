package styles

import "charm.land/lipgloss/v2"

var (
	DashboardTitle      lipgloss.Style
	DashboardItem       lipgloss.Style
	DashboardActiveItem lipgloss.Style
	DashboardContainer  lipgloss.Style
)
var (
	SSHInfoTitle      lipgloss.Style
	SSHInfoLabel      lipgloss.Style
	SSHInfoValue      lipgloss.Style
	SSHInfoTag        lipgloss.Style
	SSHInfoContainer  lipgloss.Style
	SSHInfoEmptyState lipgloss.Style
)
var (
	HelpMenu lipgloss.Style
)

func init() {
	DashboardTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("228")).
		MarginBottom(1)
	DashboardItem = lipgloss.NewStyle().PaddingLeft(1)
	DashboardActiveItem = lipgloss.NewStyle().
		PaddingLeft(1).
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("62"))
	DashboardContainer = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)

	SSHInfoTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("225")).
		MarginBottom(1)
	SSHInfoLabel = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("110"))
	SSHInfoValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))
	SSHInfoTag = lipgloss.NewStyle().
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("24")).
		Padding(0, 1).
		MarginRight(1)
	SSHInfoContainer = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("99")).
		Padding(1, 2)
	SSHInfoEmptyState = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244")).
		Italic(true)

	HelpMenu = lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		MarginTop(1)
}
