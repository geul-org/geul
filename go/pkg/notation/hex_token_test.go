//ff:func feature=notation type=decoder control=sequence
//ff:what TestHexToken — four lowercase hex digits per word after 0x
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHexToken(t *testing.T) {
	if got := hexToken([]core.Word{0x0B86, 0x00ff}); got != "0x0b8600ff" {
		t.Fatal(got)
	}
}
