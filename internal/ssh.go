package internal

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/yorukot/termuru/internal/domain"
)

func (m Model) selectedHost() *domain.SSHHost {
	if len(m.filtered) == 0 {
		return nil
	}

	cursor := m.dashboard.Cursor()
	if cursor < 0 || cursor >= len(m.filtered) {
		return nil
	}

	return &m.filtered[cursor]
}

func sshArgsFromHost(s domain.SSHHost) []string {
	args := []string{"-t"}

	if s.Port > 0 {
		args = append(args, "-p", fmt.Sprint(s.Port))
	}
	if s.PrivateKeyPath != "" {
		args = append(args, "-i", s.PrivateKeyPath)
	}

	target := s.Host
	if s.User != "" && s.Host != "" {
		target = s.User + "@" + s.Host
	}

	args = append(args, target)
	return args
}

func (m *Model) enterSSHCard() tea.Cmd {
	selected := m.selectedHost()
	if selected == nil || selected.Host == "" {
		return nil
	}

	m.sshArgs = sshArgsFromHost(*selected)
	return tea.Quit
}

func (m Model) SSHArgs() []string {
	if len(m.sshArgs) == 0 {
		return nil
	}
	out := make([]string, len(m.sshArgs))
	copy(out, m.sshArgs)
	return out
}
