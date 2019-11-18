package log

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSimpleAdapter(t *testing.T) {
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
	a := NewConsoleAdapter(b, NewFilters(map[string][]string{"default": {"-no"}}))
	_ = a.Log(e)
	require.Equal(t, "2011-11-11T11:11:11+0000 [test:test:test:test] {\"test\":\"test\"} test\n", b.String())
}
