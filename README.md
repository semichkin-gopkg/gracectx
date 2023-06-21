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
	srv := &http.Server{
		Addr: ":8000",
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-gracectx.Background().Done()
	_ = srv.Shutdown(context.Background())
}

```