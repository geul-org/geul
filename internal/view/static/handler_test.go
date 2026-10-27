//ff:func feature=server type=handler control=iteration dimension=1
//ff:what TestHandler — each vendor script and the license notice are served; unknown files are 404
package static

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	h := Handler()
	for _, f := range []string{"cytoscape.min.js", "dagre.min.js", "cytoscape-dagre.js", "LICENSES.md"} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, Prefix+f, nil))
		if rec.Code != 200 || rec.Body.Len() == 0 {
			t.Fatalf("%s: %d", f, rec.Code)
		}
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, Prefix+"none.js", nil))
	if rec.Code != 404 {
		t.Fatalf("missing: %d", rec.Code)
	}
}
