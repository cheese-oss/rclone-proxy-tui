package service

import (
	"strings"
	"testing"

	"github.com/cheese-oss/rclone-proxy-tui/internal/state"
)

func TestUnit(t *testing.T) {
	m := &Manager{Paths: state.Paths{Dir: "/home/u/my dir"}, Exe: "/usr/local/bin/rclone-proxy-tui"}
	u := m.Unit()
	for _, want := range []string{
		`ExecStart=/usr/local/bin/rclone-proxy-tui daemon --dir "/home/u/my dir"`,
		"Restart=on-failure",
		"TimeoutStopSec=90",
	} {
		if !strings.Contains(u, want) {
			t.Errorf("unit missing %q:\n%s", want, u)
		}
	}
}
