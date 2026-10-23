//ff:func feature=resolve type=encoder control=sequence
//ff:what TestFormatVersionList — one line per entry, deprecated marked
package resolve

import (
	"strings"
	"testing"
)

func TestFormatVersionList(t *testing.T) {
	z := strings.Repeat("0", 64)
	got := formatVersionList([]VersionEntry{{Version: 1}, {Version: 2, Deprecated: true}})
	if got != "0x00000001 0x"+z+"\n0x00000002 0x"+z+" deprecated\n" {
		t.Fatalf("got %q", got)
	}
}
