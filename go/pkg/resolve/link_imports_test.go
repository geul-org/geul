//ff:func feature=resolve type=helper control=sequence
//ff:what TestLinkImports — IMPORT IDs point at the selected Nodes
package resolve

import (
	"github.com/geul-org/geul/go/pkg/inspect"
	"testing"
)

func TestLinkImports(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc(3, "C", [3]uint32{1, 2, 0}, [3]uint32{2, 1, 2}))
	a, b := &Node{Version: 5}, &Node{Version: 1}
	n := &Node{Doc: doc}
	linkImports(n, map[DocKey]*Node{sampleKey(1, 0).DocKey(): a, sampleKey(2, 0).DocKey(): b})
	if n.Imports[1] != b || n.Imports[2] != a {
		t.Fatalf("got %+v", n.Imports)
	}
}
