//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseLatest — the canonical line only: generation ≥ 1, 64 lowercase hex, single space, newline
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestParseLatest(t *testing.T) {
	h := strings.Repeat("ab", 32)
	if g, n, err := parseLatest([]byte("7 " + h + "\n")); err != nil || g != 7 || n != h {
		t.Fatal(g, n, err)
	}
	for _, s := range []string{"7 " + h, "07 " + h + "\n", "0 " + h + "\n", "7 " + strings.ToUpper(h) + "\n", "7  " + h + "\n", "x\n", "7 abc\n"} {
		if _, _, err := parseLatest([]byte(s)); !errors.Is(err, ErrCorrupt) {
			t.Fatal(s, err)
		}
	}
}
