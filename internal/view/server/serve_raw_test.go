//ff:func feature=server type=handler control=sequence
//ff:what TestServeRaw — the file's bytes under the GEUL media type; a missing file is 404
package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeRaw(t *testing.T) {
	file := filepath.Join(treeFixture(t), ".well-known", "geul", "h", "a.geul")
	rec := httptest.NewRecorder()
	serveRaw(rec, httptest.NewRequest(http.MethodGet, "/a.geul", nil), file)
	want, _ := os.ReadFile(file)
	if rec.Code != 200 || rec.Body.String() != string(want) || rec.Header().Get("Content-Type") != MediaType {
		t.Fatalf("got %d %v", rec.Code, rec.Header())
	}
	rec = httptest.NewRecorder()
	serveRaw(rec, httptest.NewRequest(http.MethodGet, "/x.geul", nil), file+".none")
	if rec.Code != 404 {
		t.Fatalf("missing: %d", rec.Code)
	}
}
