//ff:func feature=server type=handler control=sequence
//ff:what TestServeDoc — a browser gets the viewer page, a GEUL client the bytes; both vary by Accept
package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestServeDoc(t *testing.T) {
	root := treeFixture(t)
	rec := httptest.NewRecorder()
	serveDoc(rec, httptest.NewRequest(http.MethodGet, "/", nil), root, "/.well-known/geul/h/a")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "browseBtn") || rec.Header().Get("Vary") != "Accept" {
		t.Fatalf("page: %d %v", rec.Code, rec.Header())
	}
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept", MediaType)
	rec = httptest.NewRecorder()
	serveDoc(rec, r, root, "/.well-known/geul/h/a")
	want, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "h", "a.geul"))
	if rec.Code != 200 || rec.Body.String() != string(want) || rec.Header().Get("Content-Type") != MediaType || rec.Header().Get("Vary") != "Accept" {
		t.Fatalf("bytes: %d %v", rec.Code, rec.Header())
	}
}
