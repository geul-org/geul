//ff:func feature=schema type=rule control=sequence
//ff:what TestVerdictOf — Pass, Review, Fail
package schema

import (
	"testing"
)

func TestVerdictOf(t *testing.T) {
	if verdictOf(nil) != Pass || verdictOf([]Fact{{Severity: Warn}}) != Review || verdictOf([]Fact{{Severity: Warn}, {Severity: Error}}) != Fail {
		t.Fatal("wrong verdict")
	}
}
