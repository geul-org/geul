//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestHTTPFetch — 200 body, 404/410 ErrNotFound, 500 error, bad URL and transport errors
package resolve

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPFetch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/.well-known/geul/ok":
			w.Write([]byte("x"))
		case "/.well-known/geul/gone":
			w.WriteHeader(http.StatusGone)
		case "/.well-known/geul/boom":
			w.WriteHeader(http.StatusInternalServerError)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()
	h := HTTP{Base: srv.URL + "/"}
	ctx := context.Background()
	if b, err := h.Fetch(ctx, "ok"); err != nil || string(b) != "x" {
		t.Fatalf("ok: %q %v", b, err)
	}
	for _, p := range []string{"gone", "missing"} {
		if _, err := h.Fetch(ctx, p); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%s: %v", p, err)
		}
	}
	if _, err := (HTTP{Base: srv.URL, Client: srv.Client()}).Fetch(ctx, "boom"); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatalf("boom: %v", err)
	}
	if _, err := (HTTP{Base: "http://\x7f"}).Fetch(ctx, "ok"); err == nil {
		t.Fatal("bad URL: want error")
	}
	srv.Close()
	if _, err := h.Fetch(ctx, "ok"); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatalf("closed: %v", err)
	}
}
