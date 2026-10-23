//ff:func feature=schema type=rule control=sequence
//ff:what TestPrintableASCII — visible ASCII only
package schema

import (
	"testing"
)

func TestPrintableASCII(t *testing.T) {
	if !printableASCII([]byte("a-Z")) || printableASCII([]byte("a b")) || printableASCII([]byte{0x80}) {
		t.Fatal("wrong")
	}
}
