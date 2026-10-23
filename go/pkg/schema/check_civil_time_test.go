//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckCivilTime — in range passes; wrong width, negative and a full day fail V2
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckCivilTime(t *testing.T) {
	ct := func(p ...core.Word) rec { return recOf(core.Record{Kind: core.KindValue, Codec: 0x21, Payload: p}) }
	day := uint64(86400 * 1e9)
	full := ct(core.Word(day>>48), core.Word(day>>32), core.Word(day>>16), core.Word(day))
	if checkCivilTime(ct(0, 0, 0, 1)) != nil || rulesOf(checkCivilTime(ct(1))) != "V2" || rulesOf(checkCivilTime(ct(0xFFFF, 0, 0, 0))) != "V2" || rulesOf(checkCivilTime(full)) != "V2" {
		t.Fatal("wrong")
	}
}
