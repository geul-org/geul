//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestSourceForm — renders and re-encodes the SOURCE; a malformed SOURCE fails
package geulmd

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestSourceForm(t *testing.T) {
	w := toulmin.Wire{ID: 1, Span: 2}
	ws, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 3, Bytes: []byte("hi")}, w)
	r, _, _ := core.DecodeRecord(ws)
	line, enc, sw, err := sourceForm(r, w)
	if err != nil || line != `SOURCE #3 span=2 "hi"` || sw != 2 || !slices.Equal(enc, ws) {
		t.Fatalf("got %q %d %v", line, sw, err)
	}
	if _, _, _, err := sourceForm(core.Record{Kind: core.KindEdge, Codec: 0x60}, w); err == nil {
		t.Fatal("not a SOURCE: want error")
	}
}
