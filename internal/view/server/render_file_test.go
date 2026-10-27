//ff:func feature=server type=handler control=sequence
//ff:what TestRenderFile — page with title, CSP and the tree browser when asked, 500 for a missing file, 422 for a non-document, 500 when writing fails
package server

import (
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderFile(t *testing.T) {
	root := treeFixture(t)
	a := filepath.Join(root, ".well-known", "geul", "h", "a.geul")
	rec := httptest.NewRecorder()
	renderFile(rec, a, "h/a.geul", true, nil)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "<title>h/a.geul · geul</title>") || rec.Header().Get("Content-Security-Policy") != contentPolicy || !strings.Contains(rec.Body.String(), `id="browseBtn"`) {
		t.Fatalf("page: %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	renderFile(rec, filepath.Join(root, "none.geul"), "none", false, nil)
	if rec.Code != 500 {
		t.Fatalf("missing: %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	renderFile(rec, filepath.Join(root, ".well-known", "geul", "h", "bad.geul"), "bad", false, nil)
	if rec.Code != 422 || !strings.HasPrefix(rec.Body.String(), "bad: ") {
		t.Fatalf("bad: %d %q", rec.Code, rec.Body)
	}
	bw := brokenWriter{httptest.NewRecorder()}
	renderFile(bw, a, "a", false, nil)
	if bw.Code != 500 {
		t.Fatalf("broken writer: %d", bw.Code)
	}
}
