package log

import (
	"context"
	"github.com/novakit/log/labels"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Setup(Options{
		Project:  "test",
		Env:      "test",
		Hostname: "test",
		Topics:   []string{"-"},
		Console: &ConsoleOptions{
			Enabled: true,
			Topics:  []string{"-"},
		},
		File: &FileOptions{
			Enabled: true,
			Dir:     "testlog",
			Topics:  []string{"-debug"},
		},
	})
	ctx := labels.Set(context.Background(), "hello", "world")
	Info(ctx, "hello, world")
	Debug(ctx, "hello, world")
	setActiveAppenders(nil)

	tm, err := time.Parse("2006-01-02T15:04:05.000-0700", "2019-11-22T15:49:44.630+0800")
	require.NoError(t, err)
	t.Log(tm)
}
