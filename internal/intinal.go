package internal

import (
	"github.com/yorukot/termuru/internal/domain"
	"github.com/yorukot/termuru/internal/ui/dashboard"
	"github.com/yorukot/termuru/internal/ui/sshinfo"
)

func IntinalModel() Model {
	state := domain.AppState{Hosts: mockHosts()}
	filtered := make([]domain.SSHHost, len(state.Hosts))
	copy(filtered, state.Hosts)

	return Model{
		state:     state,
		filtered:  filtered,
		dashboard: dashboard.NewModel(),
		sshInfo:   sshinfo.NewModel(),
	}
}

func mockHosts() []domain.SSHHost {
	return []domain.SSHHost{
		{
			Name:           "prod-api",
			Host:           "100.64.0.4",
			Port:           22,
			User:           "root",
			PrivateKeyPath: "~/.ssh/yorukot_personal",
			Description:    "你好你好你好你好你好你好你好 測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試測試",
			Environment:    "production",
			AuthMethod:     "ed25519 key",
			LastConnected:  "2026-03-01 23:42 UTC",
			Fingerprint:    "SHA256:6FiT8wG7xV2...",
			Tags:           []string{"critical", "api", "aws"},
		},
		{
			Name:           "prod-db",
			Host:           "10.0.10.20",
			Port:           22,
			User:           "postgres",
			PrivateKeyPath: "~/.ssh/yorukot_personal",
			Description:    "Primary PostgreSQL node for billing and auth",
			Environment:    "production",
			AuthMethod:     "hardware key",
			LastConnected:  "2026-02-28 14:19 UTC",
			Fingerprint:    "SHA256:Qz9a2jK2cY7...",
			Tags:           []string{"database", "restricted", "stateful"},
		},
		{
			Name:           "staging-web",
			Host:           "10.0.20.12",
			Port:           2222,
			User:           "deploy",
			PrivateKeyPath: "~/.ssh/yorukot_personal",
			Description:    "Staging web node for release candidate verification",
			Environment:    "staging",
			AuthMethod:     "ed25519 key",
			LastConnected:  "2026-03-02 03:05 UTC",
			Fingerprint:    "SHA256:Kf7n1rZ4bM9...",
			Tags:           []string{"staging", "frontend", "canary"},
		},
		{
			Name:           "sandbox",
			Host:           "192.168.56.101",
			Port:           22,
			User:           "root",
			PrivateKeyPath: "",
			Description:    "Local VM for quick shell experiments and SSH config tests",
			Environment:    "local",
			AuthMethod:     "password",
			LastConnected:  "2026-03-02 06:00 UTC",
			Fingerprint:    "SHA256:Tt4s8nU1pA3...",
			Tags:           []string{"dev", "unsafe", "lab"},
		},
	}
}
