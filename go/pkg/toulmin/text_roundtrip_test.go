//ff:func feature=toulmin type=helper control=iteration dimension=1
//ff:what TestTextRoundTrip — TEXT records round-trip empty, odd, even and multibyte strings verbatim
package toulmin

import (
	"bytes"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestTextRoundTrip checks DecodeText(EncodeText(p)) keeps the Edge Id and the
// bytes verbatim, including a decomposed Hangul sequence that NFC would change.
func TestTextRoundTrip(t *testing.T) {
	for i, s := range []string{"", "a", "ab", "한글", "가"} {
		in := TextPacket{TID: TID(i + 1), Bytes: []byte(s)}
		ws, err := EncodeText(in, testWire)
		if err != nil {
			t.Fatal(err)
		}
		r, _, _ := core.DecodeRecord(ws)
		got, err := DecodeText(r, testWire)
		if err != nil || got.TID != in.TID || !bytes.Equal(got.Bytes, in.Bytes) {
			t.Fatalf("%q: got %+v %v", s, got, err)
		}
	}
}
