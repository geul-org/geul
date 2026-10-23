//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestParseVersionLine — plain and deprecated lines; bad field counts, marks, IDs and version 0 fail
package resolve

import (
	"strings"
	"testing"
)

func TestParseVersionLine(t *testing.T) {
	h := "0x" + strings.Repeat("0", 62) + "ff"
	e, err := parseVersionLine("0x00000002 " + h)
	if err != nil || e.Version != 2 || e.Hash[31] != 0xff || e.Deprecated {
		t.Fatalf("plain: %+v %v", e, err)
	}
	if e, err := parseVersionLine("0x00000003 " + h + " deprecated"); err != nil || !e.Deprecated {
		t.Fatalf("deprecated: %+v %v", e, err)
	}
	for _, s := range []string{"0x00000002", "0x00000002 " + h + " old", "0x0002 " + h, "0x00000002 0x00", "0x00000000 " + h} {
		if _, err := parseVersionLine(s); err == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}
