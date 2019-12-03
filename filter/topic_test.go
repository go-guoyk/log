package filter

import (
	"github.com/novakit/log/event"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter_IsTopicEnabled(t *testing.T) {
	fs := Topic([]string{"info"})
	require.True(t, fs.Check(event.Event{Topic: "info"}))
	require.False(t, fs.Check(event.Event{Topic: "debug"}))
	fs = Topic([]string{"-info"})
	require.False(t, fs.Check(event.Event{Topic: "info"}))
	require.True(t, fs.Check(event.Event{Topic: "debug"}))
	require.True(t, fs.Check(event.Event{Topic: "error"}))
}
