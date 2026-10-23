//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDecimal — too short fails; the coefficient follows the integer rule
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckDecimal(t *testing.T) {
	d := func(p ...core.Word) rec { return recOf(core.Record{Kind: core.KindValue, Codec: 0x10, Payload: p}) }
	if checkDecimal(d(0, 0, 1), true) != nil || rulesOf(checkDecimal(d(0, 0), true)) != "V3" || rulesOf(checkDecimal(d(0, 0, 0, 1), false)) != "V3" {
		t.Fatal("wrong")
	}
}
