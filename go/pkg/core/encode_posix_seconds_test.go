//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodePosixSeconds — round trip through DecodeTimeValue, including a negative count
package core

import "testing"

func TestEncodePosixSeconds(t *testing.T) {
	for _, s := range []int64{0, 1700000000, -1} {
		if tm, err := DecodeTimeValue(EncodePosixSeconds(s)); err != nil || tm.Unix() != s {
			t.Fatalf("%d: %v %v", s, tm, err)
		}
	}
}
