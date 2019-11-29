package log

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter_IsTopicEnabled(t *testing.T) {
	fs := NewFilter([]string{"info"})
	require.True(t, fs.IsTopicEnabled("info"))
	require.False(t, fs.IsTopicEnabled("debug"))
	fs = NewFilter([]string{"-info"})
	require.False(t, fs.IsTopicEnabled("info"))
	require.True(t, fs.IsTopicEnabled("debug"))
	require.True(t, fs.IsTopicEnabled("error"))
}
