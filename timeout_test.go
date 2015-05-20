package timeout

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rkgo/web"
)

func TestTimeout(t *testing.T) {
	app := web.New()
	app.Use(Timeout("10ms"))
	app.Use(func(ctx web.Context, next web.Next) {
		time.Sleep(20 * time.Millisecond)

		ctx.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, (*http.Request)(nil))

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("503 Service Unavailable expected, got: %d %s", rec.Code, http.StatusText(rec.Code))
	}
}

func TestNoTimeout(t *testing.T) {
	app := web.New()
	app.Use(Timeout("10ms"))
	app.Use(func(ctx web.Context, next web.Next) {
		time.Sleep(5 * time.Millisecond)

		ctx.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, (*http.Request)(nil))

	if rec.Code != http.StatusNoContent {
		t.Errorf("204 No Content expected, got: %d %s", rec.Code, http.StatusText(rec.Code))
	}
}
