//ff:func feature=geulmd type=helper control=sequence
//ff:what TestKindName — every record kind has its keyword
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestKindName(t *testing.T) {
	if kindName(core.KindReference) != "REFERENCE" || kindName(core.KindValue) != "VALUE" || kindName(core.KindEdge) != "EDGE" || kindName(core.KindMeta) != "META" {
		t.Fatal("kind keyword mismatch")
	}
}
