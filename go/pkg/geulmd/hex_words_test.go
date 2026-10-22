//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestHexWords — four-digit lowercase words, space separated
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHexWords(t *testing.T) {
	if got := hexWords([]core.Word{1, 0xABCD}); got != "0x0001 0xabcd" {
		t.Fatalf("got %q", got)
	}
}
