//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeSignedDecimalBig — small values match EncodeDecimalValue
package core

import (
	"math/big"
	"slices"
	"testing"
)

func TestEncodeSignedDecimalBig(t *testing.T) {
	small, _ := EncodeDecimalValue(true, -2, 120)
	if got := EncodeSignedDecimalBig(-2, big.NewInt(120)); !slices.Equal(got.Payload, small.Payload) || got.Codec != small.Codec {
		t.Fatalf("small: %04X", got.Payload)
	}
}
