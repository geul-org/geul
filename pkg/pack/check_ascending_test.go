//ff:func feature=pack type=rule control=sequence
//ff:what TestCheckAscending — strictly ascending passes, equal neighbours fail
package pack

import (
	"errors"
	"testing"
)

func TestCheckAscending(t *testing.T) {
	if checkAscending([]string{"a", "b"}) != nil || !errors.Is(checkAscending([]string{"a", "a"}), ErrCorrupt) {
		t.Fatal("ascending")
	}
}
