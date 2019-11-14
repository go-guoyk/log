package log

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetGetRemoveLabel(t *testing.T) {
	require.NotEqual(t, 0, labelsKey)

	ctx := context.Background()
	ctx0 := ctx
	require.Nil(t, GetAllLabels(ctx))
	ctx = SetLabel(ctx, "hello", "world")
	ctx1 := ctx
	require.False(t, ctx0 == ctx1)
	require.Equal(t, Labels{"hello": "world"}, ctx.Value(labelsKey))
	ctx = SetLabel(ctx, "hello1", "world1")
	ctx2 := ctx
	require.True(t, ctx2 == ctx1)
	require.Equal(t, Labels{"hello": "world", "hello1": "world1"}, ctx.Value(labelsKey))
	require.Equal(t, Labels{"hello": "world", "hello1": "world1"}, GetAllLabels(ctx))
	require.Equal(t, "world1", GetLabel(ctx, "hello1"))
	ctx = RemoveLabel(ctx, "hello")
	require.Equal(t, Labels{"hello1": "world1"}, ctx.Value(labelsKey))
	require.Equal(t, Labels{"hello1": "world1"}, GetAllLabels(ctx))
	ctx = RemoveLabels(ctx, "hello", "hello1")
	require.Equal(t, Labels{}, GetAllLabels(ctx))
	ctx = SetLabels(ctx, Labels{"hello3": "world4", "hello5": "hello6"})
	require.Equal(t, Labels{"hello3": "world4", "hello5": "hello6"}, GetAllLabels(ctx))
	ctx = RemoveAllLabels(ctx)
	require.Equal(t, Labels{}, GetAllLabels(ctx))
}
