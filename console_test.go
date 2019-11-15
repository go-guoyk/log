package log

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConsoleAdapter(t *testing.T) {
	e := Event{
		Timestamp: time.Date(2011, 11, 11, 11, 11, 11, 0, time.UTC),
		Project:   "test",
		Env:       "test",
		Hostname:  "test",
		Topic:     "test",
		Labels:    Labels{"test": "test"},
		Message:   "test",
	}
	b := &bytes.Buffer{}
	a := consoleAdapter{w: b}
	_ = a.Log(e)
	require.Equal(t, "2011-11-11T11:11:11+0000 [test:test:test:test] {\"test\":\"test\"} test\n", b.String())
}
