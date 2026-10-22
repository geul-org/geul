//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestSourceRoundTrip — SOURCE records carry the text and the declared SPAN width
package toulmin

import (
	"bytes"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestSourceRoundTrip checks EncodeSource/DecodeSource keep the bytes and report
// back each declarable SPAN width.
func TestSourceRoundTrip(t *testing.T) {
	for _, sw := range []int{1, 2, 4, 8} {
		in := TextPacket{TID: 1, Bytes: []byte("# doc\n")}
		ws, err := EncodeSource(in, Wire{ID: 1, Span: sw})
		if err != nil {
			t.Fatal(err)
		}
		r, _, _ := core.DecodeRecord(ws)
		got, gotW, err := DecodeSource(r, 1)
		if err != nil || gotW != sw || got.TID != 1 || !bytes.Equal(got.Bytes, in.Bytes) {
			t.Fatalf("span width %d: got %+v %d %v", sw, got, gotW, err)
		}
	}
}
