//ff:func feature=notation type=encoder control=sequence
//ff:what TestSplitArgs — commas split at the top level, arguments trimmed; empty text has none
package notation

import (
	"slices"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	if got := splitArgs(` 1, "a,b" , F(2, 3), <x,y>`); !slices.Equal(got, []string{"1", `"a,b"`, "F(2, 3)", "<x,y>"}) {
		t.Fatalf("%q", got)
	}
	if splitArgs("  ") != nil {
		t.Fatal("empty")
	}
}
