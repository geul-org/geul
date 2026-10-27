//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestEmbeddedValue — one Value; empty, trailing words and non-Value fail
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestEmbeddedValue(t *testing.T) {
	m := func(p []core.Word) rec { return recOf(core.Record{Kind: core.KindMeta, Codec: 0x10, Payload: p}) }
	v := enc(core.EncodeUintValue(5))
	if r, ok := embeddedValue(m(v)); !ok || r.R.Payload[0] != 5 {
		t.Fatalf("ok: %+v", r)
	}
	for _, p := range [][]core.Word{nil, append(append([]core.Word{}, v...), 1), enc(core.EncodeLocalRef(1))} {
		if _, ok := embeddedValue(m(p)); ok {
			t.Fatalf("%04X: want false", p)
		}
	}
}
