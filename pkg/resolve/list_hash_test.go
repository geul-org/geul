//ff:func feature=resolve type=rule control=sequence
//ff:what TestListHash — equals the content hash of the list document
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestListHash(t *testing.T) {
	b := listDocOf(1, VersionEntry{Version: 1})
	ws, _ := core.ReadWords(b)
	if h, _ := core.ContentHash(ws); listHash(b) != h {
		t.Fatal("mismatch")
	}
}
