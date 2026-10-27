//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckInteger — signed and unsigned minimal forms
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckInteger(t *testing.T) {
	if checkInteger(rec{}, []core.Word{0, 0x8000}, true, "p") != nil || rulesOf(checkInteger(rec{}, []core.Word{0, 1}, true, "p")) != "V3" || rulesOf(checkInteger(rec{}, []core.Word{0, 0x8000}, false, "p")) != "V3" {
		t.Fatal("wrong")
	}
}
