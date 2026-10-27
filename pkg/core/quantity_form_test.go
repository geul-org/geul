//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestQuantityForm — integer, BINARY-FLOAT, FRACTION and decimal amounts and bounds pass; a text amount, a Value unit, an UNKNOWN-VALUE unit and a text bound fail
package core

import (
	"math/big"
	"testing"
)

func TestQuantityForm(t *testing.T) {
	d := EncodeSignedDecimalBig(0, big.NewInt(1))
	n, _ := EncodeIntegerBig(true, big.NewInt(2))
	u, _ := EncodeIntegerBig(false, big.NewInt(3))
	f, _ := EncodeFraction(Fraction{Numerator: n, Denominator: u})
	bf := Record{Kind: KindValue, Codec: CodecBinaryFloat, Payload: []Word{0x3C00}}
	absent := Record{Kind: KindReference}
	text := Record{Kind: KindValue, Codec: CodecUTF8TextEven}
	for i, rs := range [][4]Record{{d, absent, d, absent}, {n, absent, u, f}, {f, absent, absent, bf}, {bf, absent, absent, absent}} {
		if err := quantityForm(rs); err != nil {
			t.Fatalf("good %d: %v", i, err)
		}
	}
	for i, rs := range [][4]Record{{text, absent, absent, absent}, {d, d, absent, absent}, {d, absent, text, absent}, {d, absent, absent, text}, {d, {Kind: KindReference, Codec: 1}, absent, absent}} {
		if quantityForm(rs) == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
