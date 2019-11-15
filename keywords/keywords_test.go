package keywords

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddClearAllKeywords(t *testing.T) {
	ctx := context.Background()
	ctx = Add(ctx, "hello", "world", "hello")
	require.Equal(t, []string{"hello", "world"}, Get(ctx))
	ctx = Add(ctx, "hello", "world", "hello")
	require.Equal(t, []string{"hello", "world"}, Get(ctx))
	ctx = RemoveAll(ctx)
	require.Nil(t, Get(ctx))
}
