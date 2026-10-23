//ff:func feature=resolve type=helper control=sequence
//ff:what TestDocNameOf — the DOC-NAME of a named document; false for an unnamed one
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
	"testing"
)

func TestDocNameOf(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc(2, "x"))
	if n, ok := docNameOf(doc); !ok || core.FormatID(n.Doc) != "0xffff000000000002" {
		t.Fatalf("named: %+v %v", n, ok)
	}
	if _, ok := docNameOf(inspect.Document{}); ok {
		t.Fatal("unnamed: want false")
	}
}
