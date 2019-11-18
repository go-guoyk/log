package log

import "context"

type (
	labelsKeyType int

	// Labels is a type alias of map[string]interface{}
	Labels map[string]interface{}
)

const labelsKey = labelsKeyType(0)

// GetLabel get log label from context
func GetLabel(ctx context.Context, key string) interface{} {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		return m[key]
	}
	return nil
}

// SetLabel set log label to context, this may change context
func SetLabel(ctx context.Context, key string, val interface{}) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		m[key] = val
	} else {
		m = Labels{key: val}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	return ctx
}

// SetLabels set multiple log labels to context, this may change context
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

// RemoveLabel remove log label from context
func RemoveLabel(ctx context.Context, key string) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		delete(m, key)
	}
	return ctx
}

// RemoveLabels remove multiple labels from context
func RemoveLabels(ctx context.Context, keys ...string) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for _, key := range keys {
			delete(m, key)
		}
	}
	return ctx
}

// GetAllLabels get all log labels from context
func GetAllLabels(ctx context.Context) Labels {
	m, _ := ctx.Value(labelsKey).(Labels)
	return m
}

// RemoveAllLabels remove all log labels from context
func RemoveAllLabels(ctx context.Context) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for key := range m {
			delete(m, key)
		}
	}
	return ctx
}
