package log

import "context"

type labelsKeyType int

const labelsKey labelsKeyType = 1

func GetLabel(ctx context.Context, key string) interface{} {
	if m, ok := ctx.Value(labelsKey).(map[string]interface{}); ok {
		return m[key]
	}
	return nil
}

func SetLabel(ctx context.Context, key string, val interface{}) context.Context {
	if m, ok := ctx.Value(labelsKey).(map[string]interface{}); ok {
		m[key] = val
	} else {
		m = map[string]interface{}{key: val}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	return ctx
}

func SetLabels(ctx context.Context, labels map[string]interface{}) context.Context {
	var m map[string]interface{}
	var ok bool
	if m, ok = ctx.Value(labelsKey).(map[string]interface{}); !ok {
		m = map[string]interface{}{}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	for key, val := range labels {
		m[key] = val
	}
	return ctx
}

func RemoveLabel(ctx context.Context, key string) context.Context {
	if m, ok := ctx.Value(labelsKey).(map[string]interface{}); ok {
		delete(m, key)
	}
	return ctx
}

func RemoveLabels(ctx context.Context, keys ...string) context.Context {
	if m, ok := ctx.Value(labelsKey).(map[string]interface{}); ok {
		for _, key := range keys {
			delete(m, key)
		}
	}
	return ctx
}

func GetAllLabels(ctx context.Context) map[string]interface{} {
	m, _ := ctx.Value(labelsKey).(map[string]interface{})
	return m
}

func RemoveAllLabels(ctx context.Context) context.Context {
	if m, ok := ctx.Value(labelsKey).(map[string]interface{}); ok {
		for key := range m {
			delete(m, key)
		}
	}
	return ctx
}
