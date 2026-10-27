//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestQuantityForm — the expected forms pass; a text amount, a Value unit, an UNKNOWN-VALUE unit and a text bound fail
package core

import (
	"math/big"
	"testing"
)

func TestQuantityForm(t *testing.T) {
	d := EncodeSignedDecimalBig(0, big.NewInt(1))
	absent := Record{Kind: KindReference}
	text := Record{Kind: KindValue, Codec: CodecUTF8TextEven}
	if err := quantityForm([4]Record{d, absent, d, absent}); err != nil {
		t.Fatal(err)
	}
	for i, rs := range [][4]Record{{text, absent, absent, absent}, {d, d, absent, absent}, {d, absent, text, absent}, {d, absent, absent, text}, {d, {Kind: KindReference, Codec: 1}, absent, absent}} {
		if quantityForm(rs) == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
