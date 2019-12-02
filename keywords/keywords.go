package keywords

import (
	"context"
	"github.com/novakit/log/labels"
)

const (
	// LabelKey keywords key in log labels
	LabelKey = "keywords"
)

// Add add keywords to log labels, duplicated keywords will be removed
func Add(ctx context.Context, keywords ...string) context.Context {
	res, _ := labels.Get(ctx, LabelKey).(map[string]bool)
	if res == nil {
		res = map[string]bool{}
	}
	for _, k := range keywords {
		res[k] = true
	}
	return labels.Set(ctx, LabelKey, res)
}

// Remove remove keywords from log labels
func Remove(ctx context.Context, keywords ...string) context.Context {
	res, _ := labels.Get(ctx, LabelKey).(map[string]bool)
	if res == nil {
		return ctx
	}
	for _, k := range keywords {
		delete(res, k)
	}
	return labels.Set(ctx, LabelKey, res)
}

// Get get keywords from log labels
func Get(ctx context.Context) []string {
	res, _ := labels.Get(ctx, LabelKey).(map[string]bool)
	if res == nil {
		return nil
	}
	out := make([]string, 0, len(res))
	for k := range res {
		out = append(out, k)
	}
	return out
}

// Clear remove all keywords from log labels
func Clear(ctx context.Context) context.Context {
	return labels.Remove(ctx, LabelKey)
}
