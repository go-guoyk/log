package topics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter_IsTopicEnabled(t *testing.T) {
	fs := New([]string{"info"})
	require.True(t, fs.Contains("info"))
	require.False(t, fs.Contains("debug"))
	fs = New([]string{"-info"})
	require.False(t, fs.Contains("info"))
	require.True(t, fs.Contains("debug"))
	require.True(t, fs.Contains("error"))
}
