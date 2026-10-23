//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestAddEnums — an ENUM lands on its raw field; record fields, missing fields and wrong widths are skipped
package codecspec

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestAddEnums(t *testing.T) {
	fields := map[uint64]*Field{1: {Width: 1, Enums: map[string]string{}}, 2: {Enums: map[string]string{}}}
	body := []core.Record{
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 5, Field: 1, Code: []core.Word{7}}),
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 6, Field: 2, Code: []core.Word{7}}),
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 7, Field: 9, Code: []core.Word{7}}),
		core.EncodeEnumDecl(core.EnumDecl{EdgeID: 8, Field: 1, Code: []core.Word{0, 8}}),
		core.EncodeUintValue(1),
	}
	addEnums(body, fields, map[uint64]string{5: "seven"})
	if len(fields[1].Enums) != 1 || fields[1].Enums["0x0007"] != "seven" || len(fields[2].Enums) != 0 {
		t.Fatalf("got %+v %+v", fields[1], fields[2])
	}
}
