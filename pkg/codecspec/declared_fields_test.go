//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestDeclaredFields — FIELDs naming the DECLARE in document order; other DECLAREs' and unknown FIELDs are left out
package codecspec

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDeclaredFields(t *testing.T) {
	body := []core.Record{
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 2, Declare: 1, Width: 1}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 3, Declare: 9, Width: 1}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 1, Declare: 1, Width: 1}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 4, Declare: 1, Width: 1}),
		core.EncodeUintValue(1),
	}
	fields := map[uint64]*Field{1: {Name: "a"}, 2: {Name: "b"}, 3: {Name: "c"}}
	fs := declaredFields(1, body, fields)
	if len(fs) != 2 || fs[0].Name != "b" || fs[1].Name != "a" {
		t.Fatalf("got %+v", fs)
	}
	if fs := declaredFields(7, body, fields); len(fs) != 0 {
		t.Fatalf("none: got %+v", fs)
	}
}
