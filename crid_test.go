package log

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateGenerateGetRemoveCrid(t *testing.T) {
	ctx := context.Background()
	ctx = SetOrGenerateCrid(ctx, "")
	crid := GetCrid(ctx)
	require.NotEmpty(t, crid)
	require.NotEqual(t, EmptyCrid, crid)

	ctx = context.Background()
	ctx = SetOrGenerateCrid(ctx, "111")
	crid = GetCrid(ctx)
	require.Equal(t, "111", crid)

	ctx = RemoveCrid(ctx)
	crid = GetCrid(ctx)
	require.Equal(t, EmptyCrid, crid)
}
