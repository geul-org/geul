//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeWith — a counted REF list decodes; a missing count or an oversized count fails
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeWith(t *testing.T) {
	pos := 0
	if got, err := decodeWith([]core.Word{2, 5, 6}, &pos, testWire); err != nil || len(got) != 2 || got[1] != 6 {
		t.Fatalf("got %v %v", got, err)
	}
	pos = 0
	if _, err := decodeWith(nil, &pos, testWire); err == nil {
		t.Fatal("missing count: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{3, 5}, &pos, testWire); err == nil {
		t.Fatal("count past payload: want error")
	}
}
