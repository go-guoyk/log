package labels

import "context"

type (
	labelsKeyType int

	// Labels is a type alias of map[string]interface{}
	Labels map[string]interface{}
)

const labelsKey = labelsKeyType(0)

// Get get log label from context
func Get(ctx context.Context, key string) interface{} {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		return m[key]
	}
	return nil
}

// Set set log label to context, this may change context
func Set(ctx context.Context, key string, val interface{}) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		m[key] = val
	} else {
		m = Labels{key: val}
		ctx = context.WithValue(ctx, labelsKey, m)
	}
	return ctx
}

// Remove remove multiple labels from context
func Remove(ctx context.Context, keys ...string) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for _, key := range keys {
			delete(m, key)
		}
	}
	return ctx
}

// SetAll set multiple log labels to context, this may change context
func SetAll(ctx context.Context, labels Labels) context.Context {
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

// GetAll get all log labels from context
func GetAll(ctx context.Context) Labels {
	m, _ := ctx.Value(labelsKey).(Labels)
	return m
}

// Clear remove all log labels from context
func Clear(ctx context.Context) context.Context {
	if m, ok := ctx.Value(labelsKey).(Labels); ok {
		for key := range m {
			delete(m, key)
		}
	}
	return ctx
}
