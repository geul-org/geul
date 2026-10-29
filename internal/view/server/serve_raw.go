//ff:func feature=server type=handler control=sequence
//ff:what serveRaw — answer with a document's bytes under the GEUL media type
package server

import (
	"io/fs"
	"net/http"
)

// serveRaw answers with the bytes of the .geul file name in t under the GEUL
// media type (Range and conditional requests as http.ServeFileFS does).
func serveRaw(w http.ResponseWriter, r *http.Request, t fs.FS, name string) {
	w.Header().Set("Content-Type", MediaType)
	http.ServeFileFS(w, r, t, name)
}
