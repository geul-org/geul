//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParsePPath — versions.geul is slot 0; version 0, short IDs and other names are named
package pack

import (
	"strings"
	"testing"
)

func TestParsePPath(t *testing.T) {
	if r := parsePPath(strings.Split("0x0000/0x0000000000000008/versions.geul", "/")); r.space != spaceP || r.slot != 0 || r.doc != 8 {
		t.Fatal(r)
	}
	if r := parsePPath(strings.Split("0x0000/0x0000000000000008/0x0000000a.geul", "/")); r.space != spaceP || r.slot != 10 {
		t.Fatal(r)
	}
	for _, p := range []string{"0x0000/0x08/versions.geul", "0x0000/0x0000000000000008/0x00000000.geul",
		"0x0000/0x0000000000000008/0x00000001.txt", "0x0000/0x0000000000000008/x.geul"} {
		if r := parsePPath(strings.Split(p, "/")); r.space != spaceNamed {
			t.Fatal(p)
		}
	}
}
