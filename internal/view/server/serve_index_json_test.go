//ff:func feature=server type=handler control=sequence
//ff:what TestServeIndexJSON — the listing as JSON with path, parent and entries; an unreadable directory is 500
package server

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestServeIndexJSON(t *testing.T) {
	root := treeFixture(t)
	rec := httptest.NewRecorder()
	serveIndexJSON(rec, treeOf(root), "h", "/.well-known/geul/h/", nil)
	var p indexPage
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil || rec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("decode: %v %s", err, rec.Body)
	}
	if p.Path != "/.well-known/geul/h/" || p.Up != "/.well-known/geul/" || len(p.Entries) != 2 || p.Entries[0].URL != "/.well-known/geul/h/a" {
		t.Fatalf("got %+v", p)
	}
	rec = httptest.NewRecorder()
	serveIndexJSON(rec, treeOf(root), "none", "/x/", nil)
	if rec.Code != 500 {
		t.Fatalf("missing: %d", rec.Code)
	}
}
