package log

import "context"

type (
	labelsKeyType int

	Labels map[string]interface{}
)

const labelsKey = labelsKeyType(0)

func GetLabel(ctx context.Context, key string) interface{} {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		return m[key]
	}
	return nil
}

func SetLabel(ctx context.Context, key string, val interface{}) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		m[key] = val
	} else {
		m = Labels{key: val}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	return ctx
}

func SetLabels(ctx context.Context, labels Labels) context.Context {
	var m Labels
	var ok bool
	if m, ok = ctx.Value(labelsKey).(Labels); !ok {
		m = Labels{}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	for key, val := range labels {
		m[key] = val
	}
	return ctx
}

func RemoveLabel(ctx context.Context, key string) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		delete(m, key)
	}
	return ctx
}

func RemoveLabels(ctx context.Context, keys ...string) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for _, key := range keys {
			delete(m, key)
		}
	}
	return ctx
}

func GetAllLabels(ctx context.Context) Labels {
	m, _ := ctx.Value(labelsKey).(Labels)
	return m
}

func RemoveAllLabels(ctx context.Context) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for key := range m {
			delete(m, key)
		}
	}
	return ctx
}
