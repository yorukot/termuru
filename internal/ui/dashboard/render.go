package dashboard

import (
	"fmt"
	"strings"

	"github.com/yorukot/termuru/internal/domain"
	"github.com/yorukot/termuru/internal/icon"
	"github.com/yorukot/termuru/internal/styles"
	"github.com/yorukot/termuru/pkg/text"
)

func (m Model) Render(hosts []domain.SSHHost, termWidth, termHeight int) string {
	var b strings.Builder

	// Title Render
	b.WriteString(styles.DashboardTitle.Render("Termuru SSH Dashboard"))
	b.WriteString("\n")
	b.WriteString(renderSearchBar(m))
	b.WriteString("\n")

	// Content Render
	if len(hosts) == 0 {
		b.WriteString("No SSH host found.\n")
	}

	// Calculate available line width for host entries
	availableLineWidth := 0
	availableLineWidth = max(termWidth-styles.DashboardContainer.GetHorizontalFrameSize(), 1)

	start, end := m.Window(len(hosts), termHeight)
	for i := start; i < end; i++ {
		host := hosts[i]
		prefix := "  "
		style := styles.DashboardItem
		if i == m.cursor {
			prefix = icon.Cursor + " "
			style = styles.DashboardActiveItem
		}

		line := fmt.Sprintf(
			"%s%-14s %s@%s:%d  %s",
			prefix,
			host.Name,
			host.User,
			host.Host,
			host.Port,
			host.Description,
		)

		renderStyle := style
		contentWidth := availableLineWidth - style.GetHorizontalFrameSize()
		contentWidth = max(contentWidth, 1)
		line = text.TruncateText(line, contentWidth, "...")
		renderStyle = style.Width(availableLineWidth)

		b.WriteString(renderStyle.Render(line))
		b.WriteString("\n")
	}

	// Help Bar Render
	renderHelpBar(&b)

	// Container Render
	content := b.String()
	if termWidth > 0 && termHeight > 0 {
		return styles.DashboardContainer.
			Width(termWidth).
			Height(termHeight).
			Render(content)
	}

	return content
}

func renderSearchBar(m Model) string {
	return fmt.Sprintf("%s", icon.Search+" "+m.searchBar.View())
}

func renderHelpBar(b *strings.Builder) {
	helpText := "/: search  |  esc/ctrl+c: cancel input  |  j/k/up/down or wheel: move  |  enter: connect  |  q: quit"
	helpStyle := styles.HelpMenu
	b.WriteString(helpStyle.Render(helpText))
}
