//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestQuantifierName — every Quantifier name and reserved codes under a real Codec; nothing under Codec 0000, whose low bits are a slot marker number
package core

import "testing"

func TestQuantifierName(t *testing.T) {
	want := []string{"Identified", "Generic", "Free-choice", "Universal", "Variable", "Reserved", "Reserved", "Reserved"}
	for q, name := range want {
		if got := QuantifierName(RefLocal<<3 | uint8(q)); got != name {
			t.Fatalf("QuantifierName(LOCAL|%d)=%q want %q", q, got, name)
		}
		if got := QuantifierName(uint8(q)); got != "" {
			t.Fatalf("Codec 0000 low bits %d: %q is not a Quantifier", q, got)
		}
	}
	if got := QuantifierName(RefExternalIssuer16Word<<3 | 4); got != "Variable" {
		t.Fatalf("highest Codec: %q", got)
	}
}
