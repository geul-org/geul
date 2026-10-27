//ff:func feature=core type=helper control=sequence
//ff:what TestDecimalRat — 2345×10^-2 is 469/20, 12×10^2 is 1200, -5×10^0 is -5
package core

import (
	"math/big"
	"testing"
)

func TestDecimalRat(t *testing.T) {
	if got := decimalRat(-2, big.NewInt(2345)).RatString(); got != "469/20" {
		t.Fatalf("23.45: %s", got)
	}
	if got := decimalRat(2, big.NewInt(12)).RatString(); got != "1200" {
		t.Fatalf("12e2: %s", got)
	}
	if got := decimalRat(0, big.NewInt(-5)).RatString(); got != "-5" {
		t.Fatalf("-5: %s", got)
	}
}
