//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestBigUintWords — 0 is one 0x0000 WORD, 0x8000 one WORD (no sign WORD), 2^64 five WORDs; each reads back
package core

import (
	"math/big"
	"testing"
)

func TestBigUintWords(t *testing.T) {
	two64, _ := new(big.Int).SetString("18446744073709551616", 10)
	cases := map[string]*big.Int{"0x0000": big.NewInt(0), "0x8000": big.NewInt(0x8000), "0x00010000000000000000": two64}
	for want, n := range cases {
		ws := bigUintWords(n)
		if FormatID(ws) != want {
			t.Fatalf("%v: got %s want %s", n, FormatID(ws), want)
		}
		if back, err := wordsToBigUint(ws); err != nil || back.Cmp(n) != 0 {
			t.Fatalf("%v: back %v %v", n, back, err)
		}
	}
}
