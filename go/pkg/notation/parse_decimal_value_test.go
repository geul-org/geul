//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseDecimalValue — big SIGNED-DECIMAL, UNSIGNED-DECIMAL; bad text, a negative unsigned and a coefficient past int64 fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseDecimalValue(t *testing.T) {
	if r, err := parseDecimalValue(core.CodecSignedDecimal, "123456789012345678901234567890"); err != nil || formatRecord(r) != "123456789012345678901234567890" {
		t.Fatal(err)
	}
	if r, err := parseDecimalValue(core.CodecUnsignedDecimal, "1.5"); err != nil || formatRecord(r) != "UNSIGNED-DECIMAL(1.5)" {
		t.Fatal(err)
	}
	for _, s := range []string{"x", "-1", "123456789012345678901234567890"} {
		if _, err := parseDecimalValue(core.CodecUnsignedDecimal, s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
