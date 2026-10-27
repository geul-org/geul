//ff:func feature=page type=helper control=sequence
//ff:what TestFlagsText — set flags by name; gate alone names nothing
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestFlagsText(t *testing.T) {
	if got := flagsText(toulmin.EdgeFlagOnce | toulmin.EdgeFlagStrike | toulmin.EdgeFlagOrdered | toulmin.EdgeFlagGate); got != "once strike ordered" {
		t.Fatalf("got %q", got)
	}
	if got := flagsText(toulmin.EdgeFlagGate); got != "" {
		t.Fatalf("gate only: %q", got)
	}
}
