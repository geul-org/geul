//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what TestApply — removal, in-place replacement and appending give the next snapshot with the diff's HEAD; bad inputs fail
package delta

import (
	"bytes"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestApply(t *testing.T) {
	head := []core.Record{core.EncodeSpecVersion(2)}
	base := encodeDoc(head, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b"), bindOf(3, 9, "c")})
	d := encodeDoc(append(head, core.EncodeDeltaBase(1), core.EncodeDeltaRemove([]uint64{1})), []core.Record{bindOf(2, 9, "B"), bindOf(4, 9, "d")})
	want := encodeDoc(head, []core.Record{bindOf(2, 9, "B"), bindOf(3, 9, "c"), bindOf(4, 9, "d")})
	if got, err := Apply(base, d); err != nil || !bytes.Equal(got, want) {
		t.Fatalf("apply: %v", err)
	}
	bad := map[string][2][]byte{
		"base not a doc":     {[]byte{1}, d},
		"diff not a doc":     {base, []byte{1}},
		"no DELTA-BASE":      {base, encodeDoc(head, nil)},
		"remove and write":   {base, encodeDoc(append(head, core.EncodeDeltaBase(1), core.EncodeDeltaRemove([]uint64{2})), []core.Record{bindOf(2, 9, "x")})},
		"remove missing":     {base, encodeDoc(append(head, core.EncodeDeltaBase(1), core.EncodeDeltaRemove([]uint64{8})), nil)},
		"non-Edge diff body": {base, encodeDoc(append(head, core.EncodeDeltaBase(1)), []core.Record{core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab"))})},
	}
	for name, c := range bad {
		if _, err := Apply(c[0], c[1]); err == nil {
			t.Fatalf("%s: want error", name)
		}
	}
}
