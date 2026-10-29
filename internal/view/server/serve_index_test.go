//ff:func feature=server type=handler control=sequence
//ff:what TestServeIndex — links with a parent link below the top, none at the top; unreadable directory and failed write are 500
package server

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeIndex(t *testing.T) {
	root := treeFixture(t)
	tr := treeOf(root)
	rec := httptest.NewRecorder()
	serveIndex(rec, tr, "h", "/.well-known/geul/h/", nil)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), `href="/.well-known/geul/"`) || !strings.Contains(rec.Body.String(), `href="/.well-known/geul/h/a"`) {
		t.Fatalf("sub: %d %s", rec.Code, rec.Body)
	}
	rec = httptest.NewRecorder()
	serveIndex(rec, tr, ".", "/.well-known/geul/", nil)
	if strings.Contains(rec.Body.String(), "../") {
		t.Fatal("top: no parent link")
	}
	rec = httptest.NewRecorder()
	serveIndex(rec, tr, "none", "/x/", nil)
	if rec.Code != 500 {
		t.Fatalf("missing: %d", rec.Code)
	}
	bw := brokenWriter{httptest.NewRecorder()}
	serveIndex(bw, tr, ".", "/.well-known/geul/", nil)
	if bw.Code != 500 {
		t.Fatalf("broken writer: %d", bw.Code)
	}
}
