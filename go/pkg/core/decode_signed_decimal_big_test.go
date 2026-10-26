//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeSignedDecimalBig — a 31-digit coefficient round-trips with its exponent; a non-SIGNED-DECIMAL record fails
package core

import (
	"math/big"
	"testing"
)

func TestDecodeSignedDecimalBig(t *testing.T) {
	c, _ := new(big.Int).SetString("-1234567890123456789012345678901", 10)
	e, back, err := DecodeSignedDecimalBig(EncodeSignedDecimalBig(-7, c))
	if err != nil || e != -7 || back.Cmp(c) != 0 {
		t.Fatalf("big: %d %v %v", e, back, err)
	}
	if _, _, err := DecodeSignedDecimalBig(Record{Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: []Word{0, 0, 1}}); err == nil {
		t.Fatal("unsigned: want error")
	}
}
