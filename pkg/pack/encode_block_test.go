//ff:func feature=pack type=encoder control=iteration dimension=1
//ff:what TestEncodeBlock — P and H blocks round-trip through decodeBlock
package pack

import (
	"reflect"
	"strings"
	"testing"
)

func TestEncodeBlock(t *testing.T) {
	p := block{kind: blkPLeaf, keys: []string{pKey(0, 8, 0), pKey(0, 8, 1)}, locs: []Loc{{Off: 64, Len: 3}, {Off: 67, Len: 2}}}
	h := block{kind: blkHInner, keys: []string{strings.Repeat("a", 32), strings.Repeat("b", 32)}, locs: []Loc{{Obj: 1, Off: 64, Len: 9, Enc: 1}, {Obj: 1, Off: 73, Len: 9, Enc: 1}}}
	for _, b := range []block{p, h} {
		got, err := decodeBlock(encodeBlock(b))
		if err != nil || !reflect.DeepEqual(got, b) {
			t.Fatal(got, err)
		}
	}
}
