//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestTailWidth — fixed fields after the rest field sum; no rest is 0; two rests or a record after the rest fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
)

func TestTailWidth(t *testing.T) {
	rest := codecspec.Field{Width: 0xFFFF}
	for want, fs := range map[int][]codecspec.Field{
		6: {{}, rest, {Width: 4}, {Width: 2}},
		0: {{Width: 3}, {}},
	} {
		if n, ok := tailWidth(fs); !ok || n != want {
			t.Fatalf("%v: %d %v", fs, n, ok)
		}
	}
	for i, fs := range [][]codecspec.Field{{rest, rest}, {rest, {}}} {
		if _, ok := tailWidth(fs); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
