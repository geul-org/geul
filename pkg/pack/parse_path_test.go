//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParsePath — P paths, canonical h/ paths and everything else named (§3.4)
package pack

import (
	"strings"
	"testing"
)

func TestParsePath(t *testing.T) {
	r := parsePath("0x00010000/0x510000000000002a/0x00000001.geul")
	if r.space != spaceP || r.issuer != "0x00010000" || r.doc != 0x510000000000002a || r.slot != 1 {
		t.Fatal(r)
	}
	h := "h/0x" + strings.Repeat("0a", 32) + ".geul"
	if r := parsePath(h); r.space != spaceH || len(r.hash) != 32 || r.hash[0] != 0x0a {
		t.Fatal(r)
	}
	for _, p := range []string{"issuers.geul", "0x0000/documents.geul", "h/0x" + strings.Repeat("0A", 32) + ".geul",
		"h/0x" + strings.Repeat("0a", 32) + ".gel", "0xZZ/0x0000000000000001/versions.geul"} {
		if parsePath(p).space != spaceNamed {
			t.Fatal(p)
		}
	}
}
