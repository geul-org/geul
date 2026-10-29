//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseContentRange — `bytes a-b/total` (total may be *); other forms fail
package pack

import "testing"

func TestParseContentRange(t *testing.T) {
	if a, b, ok := parseContentRange("bytes 5-9/100"); !ok || a != 5 || b != 9 {
		t.Fatal(a, b)
	}
	if _, _, ok := parseContentRange("bytes 5-9/*"); !ok {
		t.Fatal("*")
	}
	for _, s := range []string{"", "items 5-9/1", "bytes 5/1", "bytes x-9/1", "bytes 5-y/1", "bytes 5-9"} {
		if _, _, ok := parseContentRange(s); ok {
			t.Fatal(s)
		}
	}
}
