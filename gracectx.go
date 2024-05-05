package gracectx

import (
	"context"
	"os/signal"
	"syscall"
)

func From(parent context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
}

func New() (context.Context, context.CancelFunc) {
	return From(context.Background())
}
