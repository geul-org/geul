//ff:func feature=pack type=helper control=sequence
//ff:what TestMissing — fs.ErrNotExist becomes resolve.ErrNotFound, other errors pass through
package pack

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestMissing(t *testing.T) {
	if !errors.Is(missing(fs.ErrNotExist, "x"), resolve.ErrNotFound) {
		t.Fatal("not exist")
	}
	other := errors.New("other")
	if missing(other, "x") != other {
		t.Fatal("other")
	}
}
