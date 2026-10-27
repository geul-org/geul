//ff:func feature=page type=helper control=sequence
//ff:what TestFuncField — a package function links its ALIAS; a local function says so
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestFuncField(t *testing.T) {
	f := funcField("func", toulmin.FuncRef{Alias: 2, Verb: 0x10}, map[uint64]string{2: "격틀"})
	if f.Key != "func" || f.Text != "verb 0x0010" || f.Refs[0] != (Ref{ID: 2, Label: "격틀"}) {
		t.Fatalf("alias: %+v", f)
	}
	if f := funcField("dst", toulmin.FuncRef{}, nil); f.Text != "verb 0x0000 local" || f.Refs != nil {
		t.Fatalf("local: %+v", f)
	}
}
