//ff:type feature=resolve type=resolver
//ff:what HTTP — a Source over `<Base>/.well-known/geul/` (geul.org, a mirror, an intranet server)
package resolve

import "net/http"

// HTTP fetches from `<Base>/.well-known/geul/<path>`. Client nil uses
// http.DefaultClient.
type HTTP struct {
	Base   string
	Client *http.Client
}
