//ff:func feature=server type=handler control=sequence
//ff:what TreeHandler — the viewer over a .well-known/geul tree: pages for .geul files, listings for directories
package server

import (
	"net/http"

	"github.com/geul-org/gogeul/pkg/resolve"

	"github.com/geul-org/gogeul/internal/view/static"
)

// TreeHandler serves the archive tree t (`.well-known/geul/`) at the same paths
// the public address space uses: a .geul file opens in the viewer, a directory
// lists its entries, and / redirects to the tree. Files are read per request.
func TreeHandler(t resolve.Store) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(static.Prefix, static.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveTree(w, r, t)
	})
	return mux
}
