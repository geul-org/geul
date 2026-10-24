//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDeclare — direct and ISSUER targets round-trip; other records, a missing target, short Issuer and trailing words fail
package core

import (
	"reflect"
	"slices"
	"testing"
)

func TestDecodeDeclare(t *testing.T) {
	for _, d := range []Declare{
		{EdgeID: 1, Target: TargetWord(KindEdge, EdgePublication)},
		{EdgeID: 5, Target: TargetWord(KindValue, CodecIssuer1Word), Issuer: []Word{0x20}, Format: 0xC001},
	} {
		got, err := DecodeDeclare(EncodeDeclare(d))
		if err != nil || !reflect.DeepEqual(got, d) {
			t.Fatalf("%+v: %+v %v", d, got, err)
		}
	}
	id := localWords(1)
	bad := []Record{
		EncodeUintValue(1),
		{Kind: KindEdge, Codec: EdgeDeclare},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: id},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{TargetWord(KindEdge, CodecIssuer2Word), 1})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{0x1018, 0x00E0, 2})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{TargetWord(KindValue, CodecIssuer1Word), 0x20, 0xC001, 7})},
	}
	for i, r := range bad {
		if _, err := DecodeDeclare(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
