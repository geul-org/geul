//ff:func feature=delta type=encoder control=sequence
//ff:what TestDefine — a REFERENCE-NODE with an equal Object keeps its Edge Id and LOCAL once; a new Object gets two new IDs
package delta

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDefine(t *testing.T) {
	prev := encodeDoc(nil, []core.Record{refNodeOf(2, 3, "http://x")})
	l, _ := NewLocalIDs(prev, 0)
	uri := core.EncodeBytesValue(0x44, []byte("http://x"))
	if e, loc := l.Define(uri); e != 2 || loc != 3 {
		t.Fatalf("kept: %d %d", e, loc)
	}
	if e, loc := l.Define(uri); e != 4 || loc != 5 {
		t.Fatalf("second: %d %d", e, loc)
	}
}
