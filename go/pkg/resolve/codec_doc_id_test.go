//ff:func feature=resolve type=helper control=sequence
//ff:what TestCodecDocID — the target word sits under 0x0000 0x0000 0x0001; Reference codec 0000 is not document 0
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestCodecDocID(t *testing.T) {
	if got := codecDocID(0x1018); !slices.Equal(got, []core.Word{0, 0, 1, 0x1018}) {
		t.Fatalf("%04X", got)
	}
	if got := codecDocID(0); got[0]|got[1]|got[2]|got[3] == 0 {
		t.Fatal("document 0")
	}
}
