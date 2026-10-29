//ff:func feature=pack type=decoder control=sequence
//ff:what parseContentRange — first and last byte of `bytes a-b/total`
package pack

import (
	"strconv"
	"strings"
)

// parseContentRange reads a 206 Content-Range value.
func parseContentRange(s string) (uint64, uint64, bool) {
	r, ok := strings.CutPrefix(s, "bytes ")
	r, _, ok2 := strings.Cut(r, "/")
	a, b, ok3 := strings.Cut(r, "-")
	x, err1 := strconv.ParseUint(a, 10, 64)
	y, err2 := strconv.ParseUint(b, 10, 64)
	return x, y, ok && ok2 && ok3 && err1 == nil && err2 == nil
}
