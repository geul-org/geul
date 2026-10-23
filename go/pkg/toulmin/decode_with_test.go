//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeWith — a counted REF list decodes; missing or wrong count, oversized count and a bad REF fail
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeWith(t *testing.T) {
	pos := 0
	if got, err := decodeWith([]core.Word{0x0882, 2, 0x00C0, 5, 0x00C0, 6}, &pos); err != nil || len(got) != 2 || got[1] != 6 || pos != 6 {
		t.Fatalf("got %v %v", got, err)
	}
	pos = 0
	if _, err := decodeWith(nil, &pos); err == nil {
		t.Fatal("missing count: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0890, 0, 2}, &pos); err == nil {
		t.Fatal("non-integer count: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0882, 3, 0x00C0, 5}, &pos); err == nil {
		t.Fatal("count past payload: want error")
	}
	pos = 0
	if _, err := decodeWith([]core.Word{0x0882, 1, 0x0882, 5}, &pos); err == nil {
		t.Fatal("non-LOCAL REF: want error")
	}
}
