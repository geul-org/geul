//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestQuantifierName — every Quantifier name, reserved codes, and only the low 3 bits read
package core

import "testing"

func TestQuantifierName(t *testing.T) {
	want := []string{"Identified", "Generic", "Free-choice", "Universal", "Variable", "Reserved", "Reserved", "Reserved", "Identified"}
	for q, name := range want {
		if got := QuantifierName(uint8(q)); got != name {
			t.Fatalf("QuantifierName(%d)=%q want %q", q, got, name)
		}
	}
}
