package log

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddClearAllKeywords(t *testing.T) {
	ctx := context.Background()
	ctx = AddKeywords(ctx, "hello", "world", "hello")
	require.Equal(t, []string{"hello", "world"}, GetKeywords(ctx))
	ctx = AddKeywords(ctx, "hello", "world", "hello")
	require.Equal(t, []string{"hello", "world"}, GetKeywords(ctx))
	ctx = RemoveAllKeywords(ctx)
	require.Nil(t, GetKeywords(ctx))
}
