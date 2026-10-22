//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestFuncStr — alias reference and verb code
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestFuncStr(t *testing.T) {
	if funcStr(toulmin.FuncRef{Alias: 2, Verb: 0x7400}) != "#2/0x7400" {
		t.Fatal("func mismatch")
	}
}
