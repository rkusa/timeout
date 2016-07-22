package timeout

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rkusa/web"
)

func TestTimeout(t *testing.T) {
	app := web.New()
	app.Use(Timeout("10ms"))
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		time.Sleep(20 * time.Millisecond)

		rw.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("503 Service Unavailable expected, got: %d %s", rec.Code, http.StatusText(rec.Code))
	}
}

func TestNoTimeout(t *testing.T) {
	app := web.New()
	app.Use(Timeout("10ms"))
	app.Use(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		time.Sleep(5 * time.Millisecond)

		rw.WriteHeader(http.StatusNoContent)
	})

	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	app.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("204 No Content expected, got: %d %s", rec.Code, http.StatusText(rec.Code))
	}
}
