//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestStructuredLine — Values, profile gating, unknown refs and SOURCE width tracking
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestStructuredLine(t *testing.T) {
	w := toulmin.Wire{ID: 1, Span: 1}
	if line, _, err := structuredLine(core.EncodeUintValue(2), headCtx{}, &w); err != nil || line != "VALUE uint(2)" {
		t.Fatalf("value: %q %v", line, err)
	}
	if _, _, err := structuredLine(core.Record{Kind: core.KindValue, Codec: 0x60}, headCtx{}, &w); err == nil {
		t.Fatal("unknown Value: want error")
	}
	edge := core.Record{Kind: core.KindEdge, Codec: toulmin.CodecCaseDeclare}
	if _, _, err := structuredLine(edge, headCtx{}, &w); err == nil {
		t.Fatal("no TANGEUL profile: want error")
	}
	tg := headCtx{tangeul: true}
	if _, _, err := structuredLine(core.Record{Kind: core.KindEdge, Codec: 0x66}, tg, &w); err == nil {
		t.Fatal("unnamed codec: want error")
	}
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 1}, toulmin.Wire{ID: 1, Span: 4})
	r, _, _ := core.DecodeRecord(src)
	if _, _, err := structuredLine(r, tg, &w); err != nil || w.Span != 4 {
		t.Fatalf("source: %v span=%d", err, w.Span)
	}
	bad := core.Record{Kind: core.KindEdge, Codec: 0x7C, Issuer: []core.Word{0}, Format: toulmin.FormatSource}
	if _, _, err := structuredLine(bad, tg, &w); err == nil || w.Span != 4 {
		t.Fatalf("bad source must fail and keep the width: %v %d", err, w.Span)
	}
	if _, _, err := structuredLine(edge, tg, &w); err == nil {
		t.Fatal("malformed node: want error")
	}
}
