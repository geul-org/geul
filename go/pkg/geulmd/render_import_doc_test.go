//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestRenderImportDoc — with and without a pinned version
package geulmd

import (
	"strings"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

func TestRenderImportDoc(t *testing.T) {
	v := inspect.ImportDoc{ID: 3, Issuer: []core.Word{1}, Doc: []core.Word{0, 0, 0, 2}}
	if got := renderImportDoc(v); got != "IMPORT-DOC #3 0x0001 0x0000000000000002" {
		t.Fatalf("plain: %q", got)
	}
	v.Version = 0x10002
	if got := renderImportDoc(v); !strings.HasSuffix(got, " v=0x00010002") {
		t.Fatalf("pinned: %q", got)
	}
}
