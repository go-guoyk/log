package crid

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateGenerateGetRemoveCrid(t *testing.T) {
	ctx := context.Background()
	ctx = SetOrGenerate(ctx, "")
	crid := Get(ctx)
	require.NotEmpty(t, crid)
	require.NotEqual(t, Empty, crid)

	ctx = context.Background()
	ctx = SetOrGenerate(ctx, "111")
	crid = Get(ctx)
	require.Equal(t, "111", crid)

	ctx = Remove(ctx)
	crid = Get(ctx)
	require.Equal(t, Empty, crid)
}
