//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeEdge — DecodeEdge rejects non-LINK records, truncation, reserved flags, F_ONCE on DO-ONCE, bad gates and leftovers
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestDecodeEdge feeds DecodeEdge malformed LINK records.
func TestDecodeEdge(t *testing.T) {
	rec := func(c uint8, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindEdge, Codec: c, Payload: p}
	}
	bad := []core.Record{
		{Kind: core.KindMeta},
		rec(CodecCaseDeclare),
		rec(CodecAttacks),
		rec(CodecAttacks, 0x0088, 1),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2),
		rec(CodecDo, 0x0088, 1, 0x0088, 2, 0x0088, 3),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2, 0x0088, 3),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2, 0x0088, 3, 0x0010, 0, 0),
		rec(CodecDoOnce, 0x0088, 1, 0x0088, 2, 0x0088, 3, 4, 1, 0, 0),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2, 0x0088, 3, 2, 0x0001),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2, 0x0088, 3, 0),
		rec(CodecAttacks, 0x0088, 1, 0x0088, 2, 0x0088, 3, 0, 0, 0, 0),
	}
	for i, r := range bad {
		if _, err := DecodeEdge(r, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
