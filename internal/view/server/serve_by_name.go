//ff:func feature=server type=handler control=sequence
//ff:what serveByName — redirect an Issuer or document name path to its ID path; 404 when no segment is a known name
package server

import (
	"net/http"
	"strings"
)

// serveByName answers a tree path that does not exist by reading its first
// segments as an Issuer name and a document name (namePath). A known name
// redirects (302) to the same path under the IDs, keeping a trailing / and the
// query. Anything else is 404.
func serveByName(w http.ResponseWriter, r *http.Request, root, p string) {
	to, ok := namePath(root, p)
	if !ok {
		http.NotFound(w, r)
		return
	}
	if strings.HasSuffix(r.URL.Path, "/") {
		to += "/"
	}
	if r.URL.RawQuery != "" {
		to += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, to, http.StatusFound)
}
