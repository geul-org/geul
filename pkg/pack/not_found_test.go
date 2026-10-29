//ff:func feature=pack type=helper control=sequence
//ff:what TestNotFound — wraps resolve.ErrNotFound with the path
package pack

import (
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestNotFound(t *testing.T) {
	err := notFound("a/b")
	if !errors.Is(err, resolve.ErrNotFound) || !strings.Contains(err.Error(), "a/b") {
		t.Fatal(err)
	}
}
