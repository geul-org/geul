//ff:func feature=schema type=rule control=sequence
//ff:what TestFactAt — fields land where they belong
package schema

import (
	"testing"
)

func TestFactAt(t *testing.T) {
	f := factAt("R2", Warn, rec{Index: 4, Word: 9}, "p", "e", "a")
	if f != (Fact{Rule: "R2", Severity: Warn, Where: Where{Word: 9, Record: 4, Part: "p"}, Expected: "e", Actual: "a"}) {
		t.Fatalf("got %+v", f)
	}
}
