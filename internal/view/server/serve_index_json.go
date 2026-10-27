//ff:func feature=server type=handler control=sequence
//ff:what serveIndexJSON — a directory of the tree as JSON for the tree browser
package server

import (
	"encoding/json"
	"net/http"
)

// serveIndexJSON answers the tree browser with the listing of dir (URL path
// urlPath) as {"Path", "Up", "Entries": [{"Name", "URL", "ID"}]}. A directory
// name ends in /; ID is set when Name is an Issuer's bound name.
func serveIndexJSON(w http.ResponseWriter, dir, urlPath string, labels map[string]string) {
	p, err := indexPageOf(dir, urlPath, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
