//ff:func feature=server type=handler control=sequence
//ff:what TestServeRaw — the file's bytes under the GEUL media type; a missing file is 404
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeRaw(t *testing.T) {
	tr, file := treeOf(treeFixture(t)), "h/a.geul"
	rec := httptest.NewRecorder()
	serveRaw(rec, httptest.NewRequest(http.MethodGet, "/a.geul", nil), tr, file)
	want, _ := tr.ReadFile(file)
	if rec.Code != 200 || rec.Body.String() != string(want) || rec.Header().Get("Content-Type") != MediaType {
		t.Fatalf("got %d %v", rec.Code, rec.Header())
	}
	rec = httptest.NewRecorder()
	serveRaw(rec, httptest.NewRequest(http.MethodGet, "/x.geul", nil), tr, file+".none")
	if rec.Code != 404 {
		t.Fatalf("missing: %d", rec.Code)
	}
}
