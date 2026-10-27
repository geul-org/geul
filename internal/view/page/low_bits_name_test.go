//ff:func feature=page type=helper control=sequence
//ff:what TestLowBitsName — a Quantifier under a real Codec; a marker number (named or reserved) under Codec 0000
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLowBitsName(t *testing.T) {
	if got := lowBitsName(core.RefLocal<<3 | 4); got != "Variable" {
		t.Fatalf("LOCAL Variable: %q", got)
	}
	if got := lowBitsName(2); got != "표지 번호" {
		t.Fatalf("NO-VALUE: %q", got)
	}
	if got := lowBitsName(6); got != "표지 번호 (예약)" {
		t.Fatalf("reserved: %q", got)
	}
}
