//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestMetaTime — a GREGORIAN-DATE Meta reads; an empty, trailing or non-time payload fails with ErrVerify
package resolve

import (
	"errors"
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestMetaTime(t *testing.T) {
	v, _ := core.EncodeRecord(core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{0, 1}})
	meta := func(p []core.Word) core.Record {
		return core.Record{Kind: core.KindMeta, Codec: core.MetaEffectiveFrom, Payload: p}
	}
	if tm, err := metaTime(meta(v)); err != nil || !tm.Equal(day(1)) {
		t.Fatalf("ok: %v %v", tm, err)
	}
	u, _ := core.EncodeRecord(core.EncodeUintValue(1))
	for _, p := range [][]core.Word{nil, append(append([]core.Word{}, v...), 0), u} {
		if _, err := metaTime(meta(p)); !errors.Is(err, ErrVerify) {
			t.Fatalf("%04X: %v", p, err)
		}
	}
}
