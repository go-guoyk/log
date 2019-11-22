package log

import (
	"encoding/json"
	"github.com/creasty/defaults"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptions(t *testing.T) {
	s, _ := json.Marshal(FileOptions{
		Enabled: true,
		Dir:     "log",
		Topics: []string{
			"-debug",
		},
	})
	t.Logf("%q", s)
	s, _ = json.Marshal(ConsoleOptions{
		Enabled: true,
		Topics: []string{
			"-debug",
		},
	})
	t.Logf("%q", s)

	o := Options{}
	err := defaults.Set(&o)
	require.NoError(t, err)
	s, _ = json.Marshal(o)
	t.Logf("%q", s)
}
