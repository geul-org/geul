//ff:type feature=pack type=resolver
//ff:what HTTPStore — a Store over `<Base>/<name>` by GET and one-range Range requests (206 and Content-Range checked)
package pack

import "net/http"

// HTTPStore reads objects from Base (no trailing slash needed). Client nil
// uses http.DefaultClient. Every ReadRange is one request with one range.
type HTTPStore struct {
	Base   string
	Client *http.Client
}
