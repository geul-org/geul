//ff:func feature=server type=handler control=sequence
//ff:what serveIndexJSON — a directory of the tree as JSON for the tree browser
package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
)

// serveIndexJSON answers the tree browser with the listing of dir (URL path
// urlPath) as {"Path", "Up", "Entries": [{"Name", "URL", "ID"}]}. A directory
// name ends in /; ID is set when Name is an Issuer's bound name.
func serveIndexJSON(w http.ResponseWriter, t fs.FS, dir, urlPath string, labels map[string]string) {
	p, err := indexPageOf(t, dir, urlPath, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
