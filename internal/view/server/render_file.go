//ff:func feature=server type=handler control=sequence
//ff:what renderFile — read, interpret and render one .geul file as the viewer page, with or without the tree browser
package server

import (
	"io/fs"
	"net/http"

	"github.com/geul-org/gogeul/internal/view/page"
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// renderFile answers with the viewer page of the file path in t under title,
// laid out by its own codec specs and spec; browse turns on the tree browser
// (tree mode only). An
// unreadable file is 500 and a file that is not one GEUL document is 422, each
// with the error as plain text.
func renderFile(w http.ResponseWriter, t fs.FS, path, title string, browse bool, spec codecspec.Spec) {
	b, err := fs.ReadFile(t, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d, err := inspect.DecodeWith(b, spec)
	if err != nil {
		http.Error(w, title+": "+err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", contentPolicy)
	p := page.Build(d, title)
	p.Browse = browse
	if err := page.Render(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
