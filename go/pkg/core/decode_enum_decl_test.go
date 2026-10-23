//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeEnumDecl — round trip; other records, a bad Edge Id, a bad field and a missing code fail
package core

import (
	"reflect"
	"slices"
	"testing"
)

func TestDecodeEnumDecl(t *testing.T) {
	e := EnumDecl{EdgeID: 6, Field: 5, Code: []Word{1}}
	if got, err := DecodeEnumDecl(EncodeEnumDecl(e)); err != nil || !reflect.DeepEqual(got, e) {
		t.Fatalf("got %+v %v", got, err)
	}
	id := localWords(6)
	for i, r := range []Record{
		EncodeUintValue(1),
		{Kind: KindEdge, Codec: EdgeEnum},
		{Kind: KindEdge, Codec: EdgeEnum, Payload: slices.Concat(id, []Word{1})},
		{Kind: KindEdge, Codec: EdgeEnum, Payload: slices.Concat(id, localWords(5))},
	} {
		if _, err := DecodeEnumDecl(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
