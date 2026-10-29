//ff:type feature=pack type=model
//ff:what counter — an http.Handler that counts GET and Range requests before serving files
package pack

import (
	"net/http"
	"sync"
)

// counter wraps a handler and counts requests.
type counter struct {
	h            http.Handler
	mu           sync.Mutex
	gets, ranges int
}
