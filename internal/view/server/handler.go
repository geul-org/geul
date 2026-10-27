//ff:func feature=server type=handler control=sequence
//ff:what Handler — the viewer HTTP handler for one .geul file: the page at / and the embedded scripts under static.Prefix
package server

import (
	"net/http"

	"github.com/geul-org/gogeul/internal/view/static"
)

// Handler returns the handler that serves the viewer page of the .geul file at
// path and the embedded scripts it loads. The file is read on every request,
// so a reload shows its current state.
func Handler(path string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(static.Prefix, static.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		servePage(w, r, path)
	})
	return mux
}
