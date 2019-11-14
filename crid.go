package log

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

const (
	CridKey = "_c"

	EmptyCrid = "-"
)

func SetOrGenerateCrid(ctx context.Context, crid string) context.Context {
	if len(crid) == 0 {
		return GenerateCrid(ctx)
	} else {
		return SetLabel(ctx, CridKey, crid)
	}
}

func GenerateCrid(ctx context.Context) context.Context {
	bytes := make([]byte, 8, 8)
	_, _ = rand.Read(bytes)
	return SetLabel(ctx, CridKey, hex.EncodeToString(bytes))
}

func GetCrid(ctx context.Context) string {
	crid, _ := GetLabel(ctx, CridKey).(string)
	if len(crid) == 0 {
		return EmptyCrid
	} else {
		return crid
	}
}

func RemoveCrid(ctx context.Context) context.Context {
	return RemoveLabel(ctx, CridKey)
}
