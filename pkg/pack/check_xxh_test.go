//ff:func feature=pack type=rule control=sequence
//ff:what TestCheckXXH — the low 32 bits of XXH64 must follow; wrong or missing is ErrCorrupt
package pack

import (
	"errors"
	"testing"
)

func TestCheckXXH(t *testing.T) {
	rest, err := checkXXH([]byte("abc"), []byte{0x99, 0x09, 0x77, 0xad, 1})
	if err != nil || len(rest) != 1 {
		t.Fatal(err)
	}
	if _, err := checkXXH([]byte("abc"), []byte{0x99, 0x09, 0x77, 0xae}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := checkXXH([]byte("abc"), []byte{1}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
