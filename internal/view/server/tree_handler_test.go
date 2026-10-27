//ff:func feature=server type=handler control=iteration dimension=1
//ff:what TestTreeHandler — a .geul under the tree and the embedded scripts are served
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTreeHandler(t *testing.T) {
	h := TreeHandler(treeFixture(t))
	for _, u := range []string{"/.well-known/geul/h/a.geul", "/.well-known/geul/_static/dagre.min.js"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, u, nil))
		if rec.Code != http.StatusOK {
			t.Fatalf("%s: %d", u, rec.Code)
		}
	}
}
