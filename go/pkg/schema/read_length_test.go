//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestReadLength — direct, both extensions canonical or not, truncated and too short
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestReadLength(t *testing.T) {
	type res struct {
		total, head   int
		canonical, ok bool
	}
	cases := []struct {
		ws   []core.Word
		want res
	}{
		{[]core.Word{0x0180}, res{4, 1, true, true}},
		{[]core.Word{0x0F00, 20}, res{20, 2, true, true}},
		{[]core.Word{0x0F00, 3}, res{3, 2, false, true}},
		{[]core.Word{0x0F00, 1}, res{1, 2, false, false}},
		{[]core.Word{0x0F80, 1, 0}, res{0x10000, 3, true, true}},
		{[]core.Word{0x0F80, 0, 20}, res{20, 3, false, true}},
		{[]core.Word{0x0F00}, res{}},
		{[]core.Word{0x0F80, 1}, res{}},
	}
	for i, c := range cases {
		total, head, canonical, ok := readLength(c.ws, 0)
		if (res{total, head, canonical, ok}) != c.want {
			t.Fatalf("case %d: %v %v %v %v", i, total, head, canonical, ok)
		}
	}
}
