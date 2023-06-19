# Usage

```go
package main

import (
	"context"
	"errors"
	"github.com/semichkin-gopkg/gracectx"
	"net/http"
)

func main() {
	gracefulShutdownCtx, cancel := gracectx.Background()
	defer cancel()

	srv := &http.Server{
		Addr: ":8000",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-gracefulShutdownCtx.Done()
	_ = srv.Shutdown(context.Background())
}

```