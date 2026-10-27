//ff:func feature=schema type=rule control=sequence
//ff:what TestEdgeID — a leading LOCAL Edge Id; no payload and a non-LOCAL lead fail
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestEdgeID(t *testing.T) {
	if id, ok := edgeID(recOf(edgeOf(0x00, 7))); !ok || id != 7 {
		t.Fatalf("ok: %d %v", id, ok)
	}
	if _, ok := edgeID(recOf(core.Record{Kind: core.KindEdge})); ok {
		t.Fatal("empty: want false")
	}
	if _, ok := edgeID(recOf(core.Record{Kind: core.KindEdge, Payload: enc(core.EncodeUintValue(1))})); ok {
		t.Fatal("value lead: want false")
	}
}
