package log

import (
	"encoding/json"
	"testing"
)

func TestOptions(t *testing.T) {
	s, _ := json.Marshal(FileOptions{
		Enabled: true,
		Dir:     "log",
		Topics: map[string][]string{
			"default": {"-debug"},
		},
	})
	t.Logf("%q", s)
	s, _ = json.Marshal(ConsoleOptions{
		Enabled: true,
		Topics: map[string][]string{
			"default": {"-debug"},
		},
	})
	t.Logf("%q", s)
	s, _ = json.Marshal(map[string][]string{
		"default": {"-debug"},
	})
	t.Logf("%q", s)
}
