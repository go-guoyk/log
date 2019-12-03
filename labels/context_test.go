package labels

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetGetRemoveLabel(t *testing.T) {
	require.NotEqual(t, 0, labelsKey)

	ctx := context.Background()
	ctx0 := ctx
	require.Nil(t, GetAll(ctx))
	ctx = Set(ctx, "hello", "world")
	ctx1 := ctx
	require.False(t, ctx0 == ctx1)
	require.Equal(t, Labels{"hello": "world"}, ctx.Value(labelsKey))
	ctx = Set(ctx, "hello1", "world1")
	ctx2 := ctx
	require.True(t, ctx2 == ctx1)
	require.Equal(t, Labels{"hello": "world", "hello1": "world1"}, ctx.Value(labelsKey))
	require.Equal(t, Labels{"hello": "world", "hello1": "world1"}, GetAll(ctx))
	require.Equal(t, "world1", Get(ctx, "hello1"))
	ctx = Remove(ctx, "hello")
	require.Equal(t, Labels{"hello1": "world1"}, ctx.Value(labelsKey))
	require.Equal(t, Labels{"hello1": "world1"}, GetAll(ctx))
	ctx = Remove(ctx, "hello", "hello1")
	require.Equal(t, 0, len(GetAll(ctx)))
	ctx = SetAll(ctx, Labels{"hello3": "world4", "hello5": "hello6"})
	require.Equal(t, Labels{"hello3": "world4", "hello5": "hello6"}, GetAll(ctx))
	ctx = Clear(ctx)
	require.Equal(t, 0, len(GetAll(ctx)))
}
