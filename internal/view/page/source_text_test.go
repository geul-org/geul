//ff:func feature=page type=builder control=sequence
//ff:what TestSourceText — the SOURCE text when present; false for a BODY without one
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestSourceText(t *testing.T) {
	if s, ok := sourceText(fixtureDoc(t).Body); !ok || s != "# 사안\n규칙" {
		t.Fatalf("got %q %v", s, ok)
	}
	if _, ok := sourceText([]inspect.Entry{{}}); ok {
		t.Fatal("no SOURCE: want false")
	}
}
