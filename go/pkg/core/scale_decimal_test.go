//ff:func feature=core type=decoder control=sequence
//ff:what TestScaleDecimal — exact rescaling up and down plus remainder and range errors
package core

import "testing"

func TestScaleDecimal(t *testing.T) {
	if v, err := ScaleDecimal(-1, 8, -4, 10000); err != nil || v != 8000 {
		t.Fatalf("up: %d %v", v, err)
	}
	if v, err := ScaleDecimal(-6, 500, -4, 10000); err != nil || v != 5 {
		t.Fatalf("down: %d %v", v, err)
	}
	if v, err := ScaleDecimal(1000000, 0, 0, 100); err != nil || v != 0 {
		t.Fatalf("zero: %d %v", v, err)
	}
	if _, err := ScaleDecimal(-5, 1, -4, 10000); err == nil {
		t.Fatal("extra digit must fail")
	}
	if _, err := ScaleDecimal(9, 1, 0, 100); err == nil {
		t.Fatal("overflow while scaling up must fail")
	}
	if _, err := ScaleDecimal(0, 101, 0, 100); err == nil {
		t.Fatal("value above limit must fail")
	}
}
