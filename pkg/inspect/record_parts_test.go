//ff:func feature=inspect type=decoder control=sequence
//ff:what TestRecordParts — remaining records become named Record parts (a Value ISSUER format by format name); an unnamed codec is "record"
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordParts(t *testing.T) {
	w := &walker{ws: []core.Word{0x0088, 7, 0x0830, 0x0A7C, 0x0020, 0xC001, 0x0A00, 0x0001}}
	recordParts(w)
	if len(w.out) != 3 || w.out[0].Name != "LOCAL" || !w.out[0].Record || w.out[1].Name != "record" || w.out[2].Name != "IPV4-ADDRESS" || w.pos != 8 {
		t.Fatalf("got %+v", w.out)
	}
}
