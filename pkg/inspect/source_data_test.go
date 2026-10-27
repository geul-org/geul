//ff:func feature=inspect type=decoder control=sequence
//ff:what TestSourceData — a SOURCE decodes and sets the span width; a non-SOURCE fails and keeps it
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestSourceData(t *testing.T) {
	w := toulmin.Wire{Span: 1}
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 7, Bytes: []byte("ab")}, toulmin.Wire{Span: 4})
	r, _, _ := core.DecodeRecord(src)
	data, enc, err := sourceData(r, &w)
	s, ok := data.(Source)
	if err != nil || !ok || s.TID != 7 || string(s.Bytes) != "ab" || s.SpanWidth != 4 || w.Span != 4 || len(enc) != len(src) {
		t.Fatalf("got %+v %v span=%d", data, err, w.Span)
	}
	if _, _, err := sourceData(core.Record{Kind: core.KindEdge, Codec: 0x60}, &w); err == nil || w.Span != 4 {
		t.Fatalf("non-SOURCE must fail and keep the width: %v %d", err, w.Span)
	}
}
