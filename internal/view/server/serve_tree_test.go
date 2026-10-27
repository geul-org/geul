//ff:func feature=server type=handler control=iteration dimension=1
//ff:what TestServeTree — redirect at /, listings and their JSON, documents with and without .geul, and 404 outside the tree, for missing paths and other files
package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeTree(t *testing.T) {
	root := treeFixture(t)
	cases := map[string]string{
		"/":                                    "302 text/html; charset=utf-8",
		"/.well-known/geul/":                   "200 text/html; charset=utf-8",
		"/.well-known/geul/h":                  "200 text/html; charset=utf-8",
		"/.well-known/geul/h/a":                "200 text/html; charset=utf-8",
		"/.well-known/geul/h/a.geul":           "200 " + MediaType,
		"/.well-known/geul/h/index.json":       "200 application/json",
		"/.well-known/geul/index.json":         "200 application/json",
		"/.well-known/geul/x/index.json":       "404 text/plain; charset=utf-8",
		"/outside.geul":                        "404 text/plain; charset=utf-8",
		"/.well-known/geul/h/x.geul":           "404 text/plain; charset=utf-8",
		"/.well-known/geul/h/versions":         "404 text/plain; charset=utf-8",
		"/.well-known/geul/../../outside.geul": "404 text/plain; charset=utf-8",
	}
	for u, want := range cases {
		rec := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.URL.Path = u
		serveTree(rec, r, root)
		if got := fmt.Sprintf("%d %s", rec.Code, rec.Header().Get("Content-Type")); got != want {
			t.Fatalf("%s: %q want %q", u, got, want)
		}
	}
}
