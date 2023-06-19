package gracectx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func New(parent context.Context) (context.Context, context.CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	ctx, cancel := context.WithCancel(parent)

	ch := make(chan os.Signal, 1)
	go func() {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	return ctx, cancel
}

func Background() (context.Context, context.CancelFunc) {
	return New(context.Background())
}
