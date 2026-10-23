//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestMinimalSigned — empty, one word, sign-carrying and removable leading words
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestMinimalSigned(t *testing.T) {
	ok := [][]core.Word{{0}, {0, 0x8000}, {0xFFFF, 0x7FFF}, {1, 2}}
	bad := [][]core.Word{nil, {0, 1}, {0xFFFF, 0xFFFF}}
	for _, ws := range ok {
		if !minimalSigned(ws) {
			t.Fatalf("%04X: want minimal", ws)
		}
	}
	for _, ws := range bad {
		if minimalSigned(ws) {
			t.Fatalf("%04X: want non-minimal", ws)
		}
	}
}
