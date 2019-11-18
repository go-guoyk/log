package log

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilters_IsTopicEnabled(t *testing.T) {
	fs := NewFilters(map[string][]string{
		"default": {"debug"},
		"test1":   {"-debug"},
		"test2":   {"info"},
	})
	require.True(t, fs.IsTopicEnabled("test1", "info"))
	require.True(t, fs.IsTopicEnabled("test2", "info"))
	require.True(t, fs.IsTopicEnabled("test3", "debug"))
	require.False(t, fs.IsTopicEnabled("test1", "debug"))
	require.False(t, fs.IsTopicEnabled("test2", "debug"))
	require.False(t, fs.IsTopicEnabled("test3", "info"))

	fs = NewFilters(map[string][]string{
		"test1": {"-debug"},
		"test2": {"info"},
	})
	require.False(t, fs.IsTopicEnabled("test3", "info"))
	require.False(t, fs.IsTopicEnabled("test3", "debug"))

	fs = NewFilters(map[string][]string{
		"test3": {},
	})
	require.False(t, fs.IsTopicEnabled("test3", "info"))
	require.False(t, fs.IsTopicEnabled("test3", "debug"))

	fs = nil
	require.False(t, fs.IsTopicEnabled("test3", "info"))
	require.False(t, fs.IsTopicEnabled("test3", "debug"))
}
