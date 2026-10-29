//ff:func feature=pack type=writer control=sequence
//ff:what TestSortedNamed — paths in byte order; a path twice is an error
package pack

import "testing"

func TestSortedNamed(t *testing.T) {
	got, err := sortedNamed([]streamKV{{Key: "b"}, {Key: "a"}})
	if err != nil || got[0].path != "a" || got[1].path != "b" {
		t.Fatalf("%v %v", got, err)
	}
	if _, err := sortedNamed([]streamKV{{Key: "a"}, {Key: "a"}}); err == nil {
		t.Fatal("twice: want error")
	}
}
