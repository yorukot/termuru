package internal

import (
	"github.com/yorukot/termuru/internal/domain"
	"github.com/yorukot/termuru/internal/ui/dashboard"
	"github.com/yorukot/termuru/internal/ui/sshinfo"
)

type Model struct {
	termWidth  int
	termHeight int
	state      domain.AppState
	filtered   []domain.SSHHost
	dashboard  dashboard.Model
	sshInfo    sshinfo.Model
	sshArgs    []string
}
