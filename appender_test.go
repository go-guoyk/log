package log

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiscardAppender(t *testing.T) {
	err := DiscardAppender.Log(Event{})
	require.NoError(t, err)
	err = DiscardAppender.Close()
	require.NoError(t, err)
}
