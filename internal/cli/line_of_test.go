//ff:func feature=cli type=helper control=sequence
//ff:what TestLineOf — a record of the document is laid out as there; one not in it alone
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestLineOf(t *testing.T) {
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2)}, nil)
	doc, _ := inspect.Decode(b)
	if got := lineOf(doc, core.EncodeSpecVersion(2)); got != "SPEC-VERSION 0x0002" {
		t.Fatal(got)
	}
	if got := lineOf(doc, core.EncodeNextLocalID(5)); got == "" {
		t.Fatal("not found: want its own line")
	}
}
