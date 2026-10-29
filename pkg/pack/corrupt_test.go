//ff:func feature=pack type=helper control=sequence
//ff:what TestCorrupt — wraps ErrCorrupt and names the part
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestCorrupt(t *testing.T) {
	err := corrupt("leaf")
	if !errors.Is(err, ErrCorrupt) || !strings.Contains(err.Error(), "leaf") {
		t.Fatal(err)
	}
}
