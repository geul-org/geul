//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatNumbered — prefixes per codec; a non-minimal ID is false
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestFormatNumbered(t *testing.T) {
	if s, _ := formatNumbered(core.RefWikidataPID, []core.Word{31}); s != "P31" {
		t.Fatal(s)
	}
	if _, ok := formatNumbered(core.RefLocal, []core.Word{0, 1}); ok {
		t.Fatal("non-minimal: want false")
	}
}
