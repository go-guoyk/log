package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testAdapter struct {
	t *testing.T
}

func (t2 *testAdapter) Log(t time.Time, project, topic string, labels map[string]interface{}, message string) error {
	assert.True(t2.t, time.Now().Sub(t) < time.Second && time.Now().Sub(t) > -time.Second)
	assert.Equal(t2.t, "adapter-test", project)
	assert.Equal(t2.t, "adapter-test", topic)
	assert.Equal(t2.t, map[string]interface{}{"hello": "world"}, labels)
	assert.Equal(t2.t, "hello, world", message)
	return nil
}

func TestAdapter(t *testing.T) {
	var d Adapter
	d = &testAdapter{t: t}
	SetProject("adapter-test")
	SetAdapter(d)
	defer SetAdapter(SimpleAdapter())
	ctx := SetLabel(context.Background(), "hello", "world")
	Log(ctx, "adapter-test", "hello, %s", "world")
}
