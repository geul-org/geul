//ff:func feature=pack type=rule control=sequence
//ff:what checkRange — ErrRange unless the response is 206, unencoded, with Content-Range exactly [off, off+n)
package pack

import (
	"fmt"
	"net/http"
)

// checkRange guards against servers that ignore Range (200 with the whole
// object), re-encode the body, or answer another range.
func checkRange(resp *http.Response, off, n uint64) error {
	if resp.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("%w: status %s", ErrRange, resp.Status)
	}
	if ce := resp.Header.Get("Content-Encoding"); ce != "" && ce != "identity" {
		return fmt.Errorf("%w: Content-Encoding %s", ErrRange, ce)
	}
	a, b, ok := parseContentRange(resp.Header.Get("Content-Range"))
	if !ok || a != off || b != off+n-1 {
		return fmt.Errorf("%w: Content-Range %q for bytes %d-%d", ErrRange, resp.Header.Get("Content-Range"), off, off+n-1)
	}
	return nil
}
