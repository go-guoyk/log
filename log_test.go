package log

import (
	"context"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetProject("logtest")
	ctx := SetLabel(context.Background(), "hello", "world")
	ctx = AddKeywords(ctx, "key1", "key2", "key3")
	ctx = SetOrGenerateCrid(ctx, "")
	go func() {
		Debug(ctx, "hello1")
		Debug(ctx, "hello1, %s", "world1")
	}()
	go func() {
		Info(ctx, "hello2")
		Info(ctx, "hello2, %s", "world2")
	}()
	go func() {
		Error(ctx, "hello3")
		Error(ctx, "hello3, %s", "world3")
	}()
	time.Sleep(time.Second)
}
