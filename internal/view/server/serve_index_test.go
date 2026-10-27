//ff:func feature=server type=handler control=sequence
//ff:what TestServeIndex — links with a parent link below the top, none at the top; unreadable directory and failed write are 500
package server

import (
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestServeIndex(t *testing.T) {
	root := treeFixture(t)
	top := filepath.Join(root, ".well-known", "geul")
	rec := httptest.NewRecorder()
	serveIndex(rec, filepath.Join(top, "h"), "/.well-known/geul/h/", nil)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), `href="/.well-known/geul/"`) || !strings.Contains(rec.Body.String(), `href="/.well-known/geul/h/a"`) {
		t.Fatalf("sub: %d %s", rec.Code, rec.Body)
	}
	rec = httptest.NewRecorder()
	serveIndex(rec, top, "/.well-known/geul/", nil)
	if strings.Contains(rec.Body.String(), "../") {
		t.Fatal("top: no parent link")
	}
	rec = httptest.NewRecorder()
	serveIndex(rec, filepath.Join(root, "none"), "/x/", nil)
	if rec.Code != 500 {
		t.Fatalf("missing: %d", rec.Code)
	}
	bw := brokenWriter{httptest.NewRecorder()}
	serveIndex(bw, top, "/.well-known/geul/", nil)
	if bw.Code != 500 {
		t.Fatalf("broken writer: %d", bw.Code)
	}
}
