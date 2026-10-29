//ff:func feature=pack type=handler control=sequence
//ff:what counter.ServeHTTP — count the request (Range or plain GET) and pass it on
package pack

import "net/http"

// ServeHTTP counts and serves.
func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	if r.Header.Get("Range") != "" {
		c.ranges++
	} else {
		c.gets++
	}
	c.mu.Unlock()
	c.h.ServeHTTP(w, r)
}
