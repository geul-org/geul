//ff:func feature=pack type=resolver control=sequence
//ff:what TestLiveLoc — a deletion marker is NotFound, anything else passes
package pack

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestLiveLoc(t *testing.T) {
	if l, err := liveLoc(Loc{Off: 3}, "p"); err != nil || l.Off != 3 {
		t.Fatal(err)
	}
	if _, err := liveLoc(Loc{Enc: encGone}, "p"); !errors.Is(err, resolve.ErrNotFound) {
		t.Fatal(err)
	}
}
