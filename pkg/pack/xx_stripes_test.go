//ff:func feature=pack type=helper control=sequence
//ff:what TestXxStripes — consumes whole 32-byte stripes and returns the tail
package pack

import "testing"

func TestXxStripes(t *testing.T) {
	_, rest := xxStripes(make([]byte, 70))
	if len(rest) != 6 {
		t.Fatal(len(rest))
	}
}
