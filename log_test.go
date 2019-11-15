package log

import (
	"context"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetEnv("test")
	SetProject("logtest")
	ctx := SetLabel(context.Background(), "hello", "world")
	go func() {
		Debugf(ctx, "hello1")
		Debugf(ctx, "hello1, %s", "world1")
	}()
	go func() {
		Infof(ctx, "hello2")
		Infof(ctx, "hello2, %s", "world2")
	}()
	go func() {
		Errorf(ctx, "hello3")
		Errorf(ctx, "hello3, %s", "world3")
	}()
	go func() {
		Logl(ctx, "access", Labels{"method": "GET"})
	}()
	time.Sleep(time.Second)
}
