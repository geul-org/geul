//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckWidth — allowed widths pass, others fail V2
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckWidth(t *testing.T) {
	r := recOf(core.Record{Kind: core.KindValue, Codec: 0x03, Payload: []core.Word{1, 2}})
	if checkWidth(r, 1, 2) != nil || rulesOf(checkWidth(r, 4)) != "V2" {
		t.Fatal("wrong")
	}
}
