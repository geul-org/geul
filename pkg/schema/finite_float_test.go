//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestFiniteFloat — finite binary16/32/64/128 values pass; NaN, ±Infinity of each width and a non-IEEE width fail
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFiniteFloat(t *testing.T) {
	for _, ws := range [][]core.Word{{0x3C00}, {0x3F80, 0}, {0x3FF0, 0, 0, 0}, {0x3FFF, 0, 0, 0, 0, 0, 0, 0}, {0x8000}} {
		if !finiteFloat(ws) {
			t.Fatalf("%x: want finite", ws)
		}
	}
	for _, ws := range [][]core.Word{{0x7C00}, {0xFF80, 0}, {0x7FF8, 0, 0, 1}, {0xFFFF, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0}, nil} {
		if finiteFloat(ws) {
			t.Fatalf("%x: want not finite", ws)
		}
	}
}
