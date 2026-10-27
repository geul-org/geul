//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestBigIntWords — small values match intWords; 31-digit and negative values round-trip through wordsToBigInt in minimal form
package core

import (
	"math/big"
	"slices"
	"testing"
)

func TestBigIntWords(t *testing.T) {
	for _, v := range []int64{0, 1, -1, 32767, 32768, -32768, -32769, 1 << 40, -(1 << 50)} {
		if got := bigIntWords(big.NewInt(v)); !slices.Equal(got, intWords(v)) {
			t.Fatalf("%d: %04X want %04X", v, got, intWords(v))
		}
	}
	for _, s := range []string{"1234567890123456789012345678901", "-1234567890123456789012345678901", "-170141183460469231731687303715884105728"} {
		v, _ := new(big.Int).SetString(s, 10)
		ws := bigIntWords(v)
		if back, err := wordsToBigInt(ws); err != nil || back.Cmp(v) != 0 {
			t.Fatalf("%s: %04X %v %v", s, ws, back, err)
		}
	}
}
