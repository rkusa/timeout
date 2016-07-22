// A timeout middleware that works well (but not exclusively) with
// [rkusa/web](https://github.com/rkusa/web).
//
//  app := web.New()
//  app.Use(timeout.Timeout("20ms"))
//
package timeout

import (
	"context"
	"net/http"
	"time"
)

func Timeout(d string) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		panic(err)
	}

	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r == nil {
			next(rw, r)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), duration)
		defer cancel()

		c := make(chan error, 1)

		go next(rw, r.WithContext(ctx))

		var err error
		select {
		case <-ctx.Done():
			err = ctx.Err()
		case err = <-c:
		}

		switch err {
		case nil: // do nothing
		case context.DeadlineExceeded:
			rw.WriteHeader(http.StatusServiceUnavailable)
		default:
			panic(err)
		}
	}
}
