//ff:func feature=pack type=decoder control=sequence
//ff:what TestSplitPKey — inverse of pKey
package pack

import "testing"

func TestSplitPKey(t *testing.T) {
	i, d, s := splitPKey(pKey(7, 1<<63, 0xFFFFFFFF))
	if i != 7 || d != 1<<63 || s != 0xFFFFFFFF {
		t.Fatal(i, d, s)
	}
}
