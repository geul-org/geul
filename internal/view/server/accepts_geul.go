//ff:func feature=server type=helper control=sequence
//ff:what acceptsGeul — whether the request's Accept names the GEUL media type
package server

import (
	"net/http"
	"strings"
)

// acceptsGeul reports whether r asks for application/vnd.geul. Browsers never
// send it, so they get the viewer page.
func acceptsGeul(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), MediaType)
}
