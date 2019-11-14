package log

import "context"

const (
	KeywordsKey = "_k"
)

func AddKeywords(ctx context.Context, keywords ...string) context.Context {
	res, _ := GetLabel(ctx, KeywordsKey).([]string)
outerLoop:
	for _, k := range keywords {
		for _, k2 := range res {
			if k == k2 {
				continue outerLoop
			}
		}
		res = append(res, k)
	}
	return SetLabel(ctx, KeywordsKey, res)
}

func GetKeywords(ctx context.Context) []string {
	res, _ := GetLabel(ctx, KeywordsKey).([]string)
	return res
}

func RemoveAllKeywords(ctx context.Context) context.Context {
	return RemoveLabel(ctx, KeywordsKey)
}
