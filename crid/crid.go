package crid

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/novakit/log/labels"
)

const (
	// LabelKey correlation id key in log labels
	LabelKey = "crid"

	// Empty empty crid
	Empty = "-"
)

// SetOrGenerate set crid to context log labels, if empty, generate a crid
func SetOrGenerate(ctx context.Context, crid string) context.Context {
	if len(crid) == 0 {
		return Generate(ctx)
	} else {
		return labels.Set(ctx, LabelKey, crid)
	}
}

// Generate generate a crid to context log labels
func Generate(ctx context.Context) context.Context {
	bytes := make([]byte, 8, 8)
	_, _ = rand.Read(bytes)
	return labels.Set(ctx, LabelKey, hex.EncodeToString(bytes))
}

// Get get crid from context log labels
func Get(ctx context.Context) string {
	crid, _ := labels.Get(ctx, LabelKey).(string)
	if len(crid) == 0 {
		return Empty
	} else {
		return crid
	}
}

// Remove remove crid from context log labels
func Remove(ctx context.Context) context.Context {
	return labels.Remove(ctx, LabelKey)
}
