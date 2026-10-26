//ff:func feature=notation type=encoder control=sequence
//ff:what TestParseField — record fields, ENUM names, raw words of the field's width; a wrong width and bad raw words fail
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseField(t *testing.T) {
	rank := codecspec.Field{Name: "Rank", Width: 1, Enums: map[string]string{"0x0002": "deprecated"}}
	if ws, err := parseField("DEPRECATED", rank, Context{}); err != nil || ws[0] != 2 {
		t.Fatal(err)
	}
	if ws, err := parseField("0x0001", rank, Context{}); err != nil || ws[0] != 1 {
		t.Fatal(err)
	}
	if _, err := parseField("0x00010002", rank, Context{}); err == nil {
		t.Fatal("width: want error")
	}
	if _, err := parseField("rank", rank, Context{}); err == nil {
		t.Fatal("not an ENUM name: want error")
	}
	if ws, err := parseField("Q5", codecspec.Field{}, Context{}); err != nil || len(ws) != 2 {
		t.Fatal(err)
	}
	if ws, err := parseField("0x00010002", codecspec.Field{Width: int(core.FieldRest)}, Context{}); err != nil || len(ws) != 2 {
		t.Fatal(err)
	}
}
