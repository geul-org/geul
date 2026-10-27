//ff:func feature=server type=handler control=sequence
//ff:what serveRaw — a .geul file's bytes as application/vnd.geul
package server

import "net/http"

// serveRaw answers with the bytes of the .geul file at file under the GEUL
// media type.
func serveRaw(w http.ResponseWriter, r *http.Request, file string) {
	w.Header().Set("Content-Type", MediaType)
	http.ServeFile(w, r, file)
}
