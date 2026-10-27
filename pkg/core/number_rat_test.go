//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestNumberRat — integers, decimals and fractions give exact values (2/4 = 1/2 = 0.5); broken payloads, a Reference and a non-number codec fail
package core

import (
	"math/big"
	"testing"
)

func TestNumberRat(t *testing.T) {
	two, _ := EncodeIntegerBig(true, big.NewInt(2))
	four, _ := EncodeIntegerBig(false, big.NewInt(4))
	half, _ := EncodeDecimalBig(true, -1, big.NewInt(5))
	f, _ := EncodeFraction(Fraction{Numerator: two, Denominator: four})
	for _, r := range []Record{f, half} {
		if v, err := NumberRat(r); err != nil || v.RatString() != "1/2" {
			t.Fatalf("%+v: %v %v", r, v, err)
		}
	}
	if v, err := NumberRat(four); err != nil || v.RatString() != "4" {
		t.Fatalf("4: %v %v", v, err)
	}
	for name, r := range map[string]Record{
		"bad integer": {Kind: KindValue, Codec: CodecSignedInteger, Payload: []Word{0, 1}},
		"bad decimal": {Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: []Word{0}},
		"reference":   {Kind: KindReference, Codec: CodecSignedInteger},
		"float":       {Kind: KindValue, Codec: CodecBinaryFloat, Payload: []Word{0x3C00}},
	} {
		if _, err := NumberRat(r); err == nil {
			t.Fatalf("%s: want error", name)
		}
	}
}
