//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what TestIsIssuer — 1, 2, 4 and 16-word canonical IDs only
package pack

import (
	"strings"
	"testing"
)

func TestIsIssuer(t *testing.T) {
	for _, s := range []string{"0x0000", "0x00010000", "0x0000000000000001", "0x" + strings.Repeat("ab", 32)} {
		if !isIssuer(s) {
			t.Fatal(s)
		}
	}
	for _, s := range []string{"0x000", "0x000000000000", "0x00AB", "h"} {
		if isIssuer(s) {
			t.Fatal(s)
		}
	}
}
