//ff:func feature=server type=handler control=sequence
//ff:what TestRenderFile — page with title, CSP and the tree browser when asked, 500 for a missing file, 422 for a non-document, 500 when writing fails
package server

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRenderFile(t *testing.T) {
	root := treeFixture(t)
	tr, a := treeOf(root), "h/a.geul"
	rec := httptest.NewRecorder()
	renderFile(rec, tr, a, "h/a.geul", true, nil)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "<title>h/a.geul · geul</title>") || rec.Header().Get("Content-Security-Policy") != contentPolicy || !strings.Contains(rec.Body.String(), `id="browseBtn"`) {
		t.Fatalf("page: %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	renderFile(rec, tr, "none.geul", "none", false, nil)
	if rec.Code != 500 {
		t.Fatalf("missing: %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	renderFile(rec, tr, "h/bad.geul", "bad", false, nil)
	if rec.Code != 422 || !strings.HasPrefix(rec.Body.String(), "bad: ") {
		t.Fatalf("bad: %d %q", rec.Code, rec.Body)
	}
	bw := brokenWriter{httptest.NewRecorder()}
	renderFile(bw, tr, a, "a", false, nil)
	if bw.Code != 500 {
		t.Fatalf("broken writer: %d", bw.Code)
	}
}
