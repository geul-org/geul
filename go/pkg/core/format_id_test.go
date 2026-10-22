//ff:func feature=core type=encoder control=sequence
//ff:what TestFormatID — one number whose digit count encodes the width
package core

import "testing"

func TestFormatID(t *testing.T) {
	if got := FormatID([]Word{0, 5}); got != "0x00000005" {
		t.Fatalf("got %q", got)
	}
}
