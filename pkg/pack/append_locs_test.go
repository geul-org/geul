//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendLocs — C4 objects, C5 offset codes, C6 lengths, C7 encodings
package pack

import (
	"bytes"
	"testing"
)

func TestAppendLocs(t *testing.T) {
	ls := []Loc{{Obj: 2, Off: 64, Len: 4, Enc: 1}, {Obj: 2, Off: 68, Len: 1}, {Enc: encGone}}
	want := []byte{2, 2, 0, 65, 0, 1, 4, 1, 0, 1, 0, 0xFF}
	if b := appendLocs(nil, ls); !bytes.Equal(b, want) {
		t.Fatalf("%x", b)
	}
}
