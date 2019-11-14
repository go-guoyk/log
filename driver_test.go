package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testDriver struct {
	t *testing.T
}

func (t2 *testDriver) Log(t time.Time, project, topic string, meta map[string]interface{}, message string) error {
	assert.True(t2.t, time.Now().Sub(t) < time.Second && time.Now().Sub(t) > -time.Second)
	assert.Equal(t2.t, "drivertest", project)
	assert.Equal(t2.t, "drivertest", topic)
	assert.Equal(t2.t, map[string]interface{}{"hello": "world"}, meta)
	assert.Equal(t2.t, "hello, world", message)
	return nil
}

func TestDriver(t *testing.T) {
	var d Driver
	d = &testDriver{t: t}
	SetProject("drivertest")
	SetDriver(d)
	defer SetDriver(SimpleDriver())
	ctx := SetLabel(context.Background(), "hello", "world")
	Log(ctx, "drivertest", "hello, %s", "world")
}
