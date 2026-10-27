//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what TestTripleAbout — a TRIPLE on a LOCAL gives that LOCAL and Predicate·Object·Rank words; other subjects, codecs and broken payloads are false
package delta

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTripleAbout(t *testing.T) {
	id, words, ok := tripleAbout(tripleOf(4, 1476, core.EncodeLocalRef(20), "ab"))
	want := []core.Word{0x00A8, 1476, 0x08C0, 0x6162, 0}
	if !ok || id != 20 || !slices.Equal(words, want) {
		t.Fatalf("%d %04x %v", id, words, ok)
	}
	bad := []core.Record{
		tripleOf(4, 1476, core.Record{Kind: core.KindReference, Codec: 0xA0, Payload: []core.Word{42}}, "a"),
		tripleOf(4, 1476, core.EncodeLocalRef(0), "a"),
		bindOf(4, 20, "a"),
		{Kind: core.KindEdge, Codec: edgeTriple, Payload: []core.Word{0x0088, 4, 0x00A9}},
		{Kind: core.KindEdge, Codec: edgeTriple, Payload: []core.Word{0x0088, 4, 0x00A8, 1, 0x0089}},
		{Kind: core.KindEdge, Codec: edgeTriple, Issuer: []core.Word{1}, Payload: []core.Word{0x0088, 4}},
		core.EncodeLocalRef(3),
	}
	for i, r := range bad {
		if _, _, ok := tripleAbout(r); ok {
			t.Fatalf("case %d: want false", i)
		}
	}
}
