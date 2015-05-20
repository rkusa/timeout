// A timeout middleware for [rkgo/web](https://github.com/rkgo/web)
//
//  app := web.New()
//  app.Use(timeout.Timeout("20ms"))
//
package timeout

import (
	"net/http"
	"time"

	"github.com/rkgo/web"
	"golang.org/x/net/context"
)

func Timeout(d string) web.Middleware {
	duration, err := time.ParseDuration(d)
	if err != nil {
		panic(err)
	}

	return func(ctx web.Context, next web.Next) {
		timeout, cancel := context.WithTimeout(ctx, duration)
		defer cancel()

		c := make(chan error, 1)
		ctx = ctx.Evolve(timeout)

		go func() {
			next(ctx)

			c <- nil
		}()

		var err error
		select {
		case <-ctx.Done():
			err = ctx.Err()
		case err = <-c:
		}

		switch err {
		case nil: // do nothing
		case context.DeadlineExceeded:
			ctx.WriteHeader(http.StatusServiceUnavailable)
		default:
			panic(err)
		}
	}
}
