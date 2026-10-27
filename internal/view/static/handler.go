//ff:func feature=server type=handler control=sequence
//ff:what Handler — the embedded third-party scripts under Prefix
package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed vendor
var files embed.FS

// Handler serves the embedded vendor files (graph view scripts and their
// licenses) under Prefix, so the page needs no network access.
func Handler() http.Handler {
	sub, _ := fs.Sub(files, "vendor")
	return http.StripPrefix(Prefix, http.FileServer(http.FS(sub)))
}
