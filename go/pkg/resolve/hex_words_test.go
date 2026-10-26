//ff:func feature=resolve type=helper control=sequence
//ff:what TestHexWords — bytes become big-endian words
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestHexWords(t *testing.T) {
	if got := hexWords([]byte{0, 1, 0xab, 0xcd}); !slices.Equal(got, []core.Word{1, 0xabcd}) {
		t.Fatalf("%04X", got)
	}
}
