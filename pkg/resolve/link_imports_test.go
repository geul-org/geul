//ff:func feature=resolve type=helper control=sequence
//ff:what TestLinkImports — IMPORT IDs point at the Node each import is bound to
package resolve

import (
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestLinkImports(t *testing.T) {
	doc, _ := inspect.Decode(sampleDoc(3, "C", [3]uint32{1, 2, 0}, [3]uint32{2, 1, 2}))
	a2, b := &Node{Version: 2}, &Node{Version: 1}
	sel := selection{
		nodes:  map[NodeKey]*Node{{DocKey: sampleKey(1, 0).DocKey(), Version: 2}: a2},
		latest: map[DocKey]*Node{sampleKey(2, 0).DocKey(): b},
	}
	n := &Node{Doc: doc}
	linkImports(n, sel)
	if n.Imports[1] != b || n.Imports[2] != a2 {
		t.Fatalf("got %+v", n.Imports)
	}
}
