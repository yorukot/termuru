package sshinfo

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/yorukot/termuru/internal/domain"
	"github.com/yorukot/termuru/internal/styles"
)

func (m Model) Render(host *domain.SSHHost, termWidth, termHeight int) string {
	var b strings.Builder

	b.WriteString(styles.SSHInfoTitle.Render("SSH Details"))
	b.WriteString("\n")

	if host == nil {
		b.WriteString(styles.SSHInfoEmptyState.Render("No host selected."))
		return fitContainer(styles.SSHInfoContainer, b.String(), termWidth, termHeight)
	}

	writeField(&b, "Name", host.Name)
	writeField(&b, "Address", fmt.Sprintf("%s@%s:%d", host.User, host.Host, host.Port))
	writeField(&b, "Environment", host.Environment)
	writeField(&b, "Auth", host.AuthMethod)
	writeField(&b, "Private Key", host.PrivateKeyPath)
	writeField(&b, "Last Connected", host.LastConnected)
	writeField(&b, "Fingerprint", host.Fingerprint)
	writeField(&b, "Description", host.Description)

	b.WriteString(styles.SSHInfoLabel.Render("Tags: "))
	if len(host.Tags) == 0 {
		b.WriteString(styles.SSHInfoValue.Render("none"))
	} else {
		for _, tag := range host.Tags {
			b.WriteString(styles.SSHInfoTag.Render(tag))
		}
	}
	b.WriteString("\n")

	return fitContainer(styles.SSHInfoContainer, b.String(), termWidth, termHeight)
}

func writeField(b *strings.Builder, label, value string) {
	b.WriteString(styles.SSHInfoLabel.Render(label + ": "))
	b.WriteString(styles.SSHInfoValue.Render(value))
	b.WriteString("\n")
}

func fitContainer(container lipgloss.Style, content string, termWidth, termHeight int) string {
	if termWidth > 0 && termHeight > 0 {
		return container.Width(termWidth).Height(termHeight).Render(content)
	}
	return container.Render(content)
}
