//ff:func feature=core type=encoder control=sequence
//ff:what TestTrimDecimal — trailing zeros removed and zero normalized
package core

import "testing"

func TestTrimDecimal(t *testing.T) {
	if e, c := TrimDecimal(-4, 8000); e != -1 || c != 8 {
		t.Fatalf("8000e-4: %d %d", e, c)
	}
	if e, c := TrimDecimal(-4, 0); e != 0 || c != 0 {
		t.Fatalf("0: %d %d", e, c)
	}
	if e, c := TrimDecimal(-4, -10000); e != 0 || c != -1 {
		t.Fatalf("-1.0: %d %d", e, c)
	}
}
