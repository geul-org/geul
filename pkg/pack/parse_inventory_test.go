//ff:func feature=pack type=decoder control=sequence
//ff:what TestParseInventory — a bad Issuer table and trailing bytes are ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestParseInventory(t *testing.T) {
	if _, err := parseInventory([]byte{1, 0, 0, 0, 3}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := parseInventory(append(encodeInventory(sampleInventory()), 0)); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
