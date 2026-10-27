//ff:func feature=notation type=encoder control=sequence
//ff:what TestFlushToken — a built token is appended and reset; an empty one is not
package notation

import (
	"strings"
	"testing"
)

func TestFlushToken(t *testing.T) {
	var b strings.Builder
	out := flushToken(nil, &b)
	b.WriteString("x")
	out = flushToken(out, &b)
	if len(out) != 1 || out[0] != "x" || b.Len() != 0 {
		t.Fatal(out)
	}
}
