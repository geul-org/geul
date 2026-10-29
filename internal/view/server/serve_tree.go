//ff:func feature=server type=handler control=selection
//ff:what serveTree — route a tree URL: redirect /, list directories (and their index.json), serve documents raw or as pages, resolve names, else 404
package server

import (
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// serveTree maps the cleaned URL path into the tree t, following the public
// address rules (GEUL-ORG 형식 지정). Paths outside /.well-known/geul/ are 404
// except /, which redirects there. A path with the .geul extension is the
// document's bytes; the same path without it is negotiated by Accept
// (serveDoc). A directory is listed, and its index.json is the listing as JSON
// for the tree browser. A missing path may be an Issuer name (serveByName);
// any other file is 404.
func serveTree(w http.ResponseWriter, r *http.Request, t resolve.Store) {
	p := path.Clean("/" + r.URL.Path)
	name := treeName(p)
	fi, err := fs.Stat(t, name)
	switch {
	case p == "/":
		http.Redirect(w, r, wellKnown, http.StatusFound)
	case !strings.HasPrefix(p+"/", wellKnown):
		http.NotFound(w, r)
	case path.Base(p) == indexJSON && isDir(t, path.Dir(name)):
		serveIndexJSON(w, t, path.Dir(name), path.Dir(p)+"/", treeLabels(t, path.Dir(p)))
	case err != nil && isFile(t, name+".geul"):
		serveDoc(w, r, t, p)
	case err != nil:
		serveByName(w, r, t, p)
	case fi.IsDir():
		serveIndex(w, t, name, strings.TrimSuffix(p, "/")+"/", treeLabels(t, p))
	case strings.HasSuffix(p, ".geul"):
		serveRaw(w, r, t, name)
	default:
		http.NotFound(w, r)
	}
}
