//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestFractionRat — 23.45/1256.1 is 469/25122 exactly; a non-FRACTION, a broken numerator or denominator and a zero denominator fail
package core

import (
	"math/big"
	"testing"
)

func TestFractionRat(t *testing.T) {
	sd, _ := EncodeDecimalBig(true, -2, big.NewInt(2345))
	ud, _ := EncodeDecimalBig(false, -1, big.NewInt(12561))
	f, _ := EncodeFraction(Fraction{Numerator: sd, Denominator: ud})
	if v, err := FractionRat(f); err != nil || v.RatString() != "469/25122" {
		t.Fatalf("got %v %v", v, err)
	}
	one, _ := EncodeIntegerBig(false, big.NewInt(1))
	zero, _ := EncodeIntegerBig(false, big.NewInt(0))
	loose := Record{Kind: KindValue, Codec: CodecSignedInteger, Payload: []Word{0, 1}}
	looseDen := Record{Kind: KindValue, Codec: CodecUnsignedInteger, Payload: []Word{0, 1}}
	zeroDen, _ := EncodeFraction(Fraction{Numerator: sd, Denominator: zero})
	badNum, _ := EncodeFraction(Fraction{Numerator: loose, Denominator: one})
	badDen, _ := EncodeFraction(Fraction{Numerator: sd, Denominator: looseDen})
	for i, r := range []Record{{Kind: KindValue, Codec: CodecQuantity}, badNum, badDen, zeroDen} {
		if _, err := FractionRat(r); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
