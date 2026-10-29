//ff:func feature=server type=handler control=sequence
//ff:what servePage — the viewer page of the single file at /; other paths are 404
package server

import (
	"net/http"
	"os"
	"path/filepath"
)

// contentPolicy allows the page's inline style and script and the scripts
// served from this server's static.Prefix, and the tree browser's fetches of this
// server's listings; nothing is fetched from elsewhere.
const contentPolicy = "default-src 'none'; style-src 'unsafe-inline'; script-src 'self' 'unsafe-inline'; connect-src 'self'"

// servePage answers GET / with the viewer page of the file at path. Other
// paths are 404.
func servePage(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderFile(w, os.DirFS(filepath.Dir(path)), filepath.Base(path), filepath.Base(path), false, nil)
}
