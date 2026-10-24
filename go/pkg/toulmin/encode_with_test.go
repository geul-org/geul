//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeWith — UNSIGNED-INTEGER count then one LOCAL REF per item; empty list is the count alone
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestEncodeWith(t *testing.T) {
	want := []core.Word{0x0882, 2, 0x0088, 5, 0x0088, 6}
	if got := encodeWith(nil, []TID{5, 6}); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X", got)
	}
	if got := encodeWith([]core.Word{9}, nil); !reflect.DeepEqual(got, []core.Word{9, 0x0882, 0}) {
		t.Fatalf("empty: %04X", got)
	}
}
