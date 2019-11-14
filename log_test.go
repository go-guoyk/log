package log

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	SetProject("logtest")
	ctx := SetLabel(context.Background(), "hello", "world")
	Debug(ctx, "hello1")
	Info(ctx, "hello2")
	Error(ctx, "hello3")
}
