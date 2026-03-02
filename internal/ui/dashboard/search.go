package dashboard

import (
	"fmt"
	"strings"

	"github.com/yorukot/termuru/internal/domain"
)

func (m Model) FilterHosts(hosts []domain.SSHHost) []domain.SSHHost {
	query := strings.TrimSpace(m.searchBar.Value())
	if query == "" {
		filtered := make([]domain.SSHHost, len(hosts))
		copy(filtered, hosts)
		return filtered
	}

	filtered := make([]domain.SSHHost, 0, len(hosts))
	for _, host := range hosts {
		if hostMatchesQuery(host, query) {
			filtered = append(filtered, host)
		}
	}
	return filtered
}

func hostMatchesQuery(host domain.SSHHost, query string) bool {
	loweredQuery := strings.ToLower(query)
	if strings.Contains(strings.ToLower(host.Name), loweredQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(host.Host), loweredQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(host.User), loweredQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(host.Description), loweredQuery) {
		return true
	}
	if strings.Contains(strings.ToLower(host.Environment), loweredQuery) {
		return true
	}
	if strings.Contains(fmt.Sprint(host.Port), loweredQuery) {
		return true
	}
	for _, tag := range host.Tags {
		if strings.Contains(strings.ToLower(tag), loweredQuery) {
			return true
		}
	}
	return false
}
