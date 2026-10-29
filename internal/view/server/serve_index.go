//ff:func feature=server type=handler control=sequence
//ff:what serveIndex — a directory of the tree as a page of links
package server

import (
	"io/fs"
	"net/http"
)

// serveIndex lists the directory dir whose URL path is urlPath (ending in /).
func serveIndex(w http.ResponseWriter, t fs.FS, dir, urlPath string, labels map[string]string) {
	p, err := indexPageOf(t, dir, urlPath, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", contentPolicy)
	if err := indexTemplate.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
