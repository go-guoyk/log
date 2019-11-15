package keywords

import (
	"context"
	"github.com/novakit/log"
)

const (
	// LabelKey keywords key in log labels
	LabelKey = "keywords"
)

// Add add keywords to log labels, duplicated keywords will be removed
func Add(ctx context.Context, keywords ...string) context.Context {
	res, _ := log.GetLabel(ctx, LabelKey).([]string)
outerLoop:
	for _, k := range keywords {
		for _, k2 := range res {
			if k == k2 {
				continue outerLoop
			}
		}
		res = append(res, k)
	}
	return log.SetLabel(ctx, LabelKey, res)
}

// Get get keywords from log labels
func Get(ctx context.Context) []string {
	res, _ := log.GetLabel(ctx, LabelKey).([]string)
	return res
}

// RemoveAll remove all keywords from log labels
func RemoveAll(ctx context.Context) context.Context {
	return log.RemoveLabel(ctx, LabelKey)
}
