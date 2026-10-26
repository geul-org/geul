//ff:func feature=cli type=helper control=sequence
//ff:what TestReferrer — the Edge referring to a gone ID; none when nothing refers
package cli

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/notation"
)

func TestReferrer(t *testing.T) {
	ref, _ := core.EncodeRecord(core.EncodeLocalRef(2))
	name, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("a")))
	d := draft{body: []core.Record{edgeRecord(3, notation.Edge{Codec: core.EdgeBind, Rest: append(ref, name...)})}}
	if by, ok := referrer(d, []uint64{2}); !ok || by != 3 {
		t.Fatalf("%d %v", by, ok)
	}
	if _, ok := referrer(d, []uint64{7}); ok {
		t.Fatal("none: want false")
	}
}
