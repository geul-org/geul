//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestFractionForm — integer and decimal components in their signed/unsigned codecs pass; a swapped sign, a Reference, a nested FRACTION and a slot marker fail
package core

import (
	"math/big"
	"testing"
)

func TestFractionForm(t *testing.T) {
	si, _ := EncodeIntegerBig(true, big.NewInt(2))
	ui, _ := EncodeIntegerBig(false, big.NewInt(3))
	sd, _ := EncodeDecimalBig(true, -2, big.NewInt(2345))
	ud, _ := EncodeDecimalBig(false, -1, big.NewInt(12561))
	for _, p := range [][2]Record{{si, ui}, {sd, ud}, {si, ud}, {sd, ui}} {
		if err := fractionForm(p[0], p[1]); err != nil {
			t.Fatalf("%+v: %v", p, err)
		}
	}
	nested, _ := EncodeFraction(Fraction{Numerator: si, Denominator: ui})
	for i, p := range [][2]Record{{ui, ui}, {si, si}, {{Kind: KindReference, Codec: CodecSignedInteger}, ui}, {si, {Kind: KindReference, Codec: CodecUnsignedInteger}}, {nested, ui}, {si, nested}, {si, {Kind: KindReference}}} {
		if fractionForm(p[0], p[1]) == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
