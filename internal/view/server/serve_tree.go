//ff:func feature=server type=handler control=selection
//ff:what serveTree — route a request under /.well-known/geul/ to a document (viewer or bytes), a listing or an Issuer name
package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// serveTree maps the cleaned URL path into root, following the public address
// rules (GEUL-ORG 형식 지정). Paths outside /.well-known/geul/ are 404 except /,
// which redirects there. A path with the .geul extension is the document's
// bytes; the same path without it is negotiated by Accept (serveDoc). A
// directory is listed, and its index.json is the listing as JSON for the tree
// browser. A missing path may be an Issuer name (serveByName); any other file
// is 404.
func serveTree(w http.ResponseWriter, r *http.Request, root string) {
	p := path.Clean("/" + r.URL.Path)
	file := filepath.Join(root, filepath.FromSlash(p))
	fi, err := os.Stat(file)
	switch {
	case p == "/":
		http.Redirect(w, r, wellKnown, http.StatusFound)
	case !strings.HasPrefix(p+"/", wellKnown):
		http.NotFound(w, r)
	case path.Base(p) == indexJSON && isDir(filepath.Dir(file)):
		serveIndexJSON(w, filepath.Dir(file), path.Dir(p)+"/", treeLabels(root, path.Dir(p)))
	case err != nil && isFile(file+".geul"):
		serveDoc(w, r, root, p)
	case err != nil:
		serveByName(w, r, root, p)
	case fi.IsDir():
		serveIndex(w, file, strings.TrimSuffix(p, "/")+"/", treeLabels(root, p))
	case strings.HasSuffix(p, ".geul"):
		serveRaw(w, r, file)
	default:
		http.NotFound(w, r)
	}
}
