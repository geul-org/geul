//ff:func feature=delta type=decoder control=sequence
//ff:what TestRemember — an Edge is kept by content, a REFERENCE-NODE also by its Object; a non-Edge is ignored
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRemember(t *testing.T) {
	l, _ := NewLocalIDs(nil, 0)
	l.remember(bindOf(4, 2, "a"))
	l.remember(refNodeOf(5, 6, "http://x"))
	l.remember(core.EncodeSpecVersion(2))
	if len(l.prev) != 2 || len(l.edges) != 2 || len(l.defs) != 1 {
		t.Fatalf("%+v", l)
	}
}
