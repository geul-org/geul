//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestHexID — one number whose digit count encodes the width
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHexID(t *testing.T) {
	if got := hexID([]core.Word{0, 5}); got != "0x00000005" {
		t.Fatalf("got %q", got)
	}
}
