package log

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	SetProject("logtest")
	ctx := SetLabel(context.Background(), "hello", "world")
	Debug(ctx, "hello1")
	Debug(ctx, "hello1, %s", "world1")
	Info(ctx, "hello2")
	Info(ctx, "hello2, %s", "world2")
	Error(ctx, "hello3")
	Error(ctx, "hello3, %s", "world3")
}
