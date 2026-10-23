//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckBool — 0 and 1 pass; 2 and wrong width fail V2
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckBool(t *testing.T) {
	b := func(p ...core.Word) rec { return recOf(core.Record{Kind: core.KindValue, Codec: 0, Payload: p}) }
	if checkBool(b(1)) != nil || rulesOf(checkBool(b(2))) != "V2" || rulesOf(checkBool(b(0, 1))) != "V2" {
		t.Fatal("wrong")
	}
}
