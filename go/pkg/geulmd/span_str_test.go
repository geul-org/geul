//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestSpanStr — offset plus length
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestSpanStr(t *testing.T) {
	if spanStr(toulmin.Span{Off: 3, Len: 4}) != "@3+4" {
		t.Fatal("span mismatch")
	}
}
