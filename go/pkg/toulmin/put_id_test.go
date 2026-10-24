//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestPutID — appends a LOCAL Reference record in minimal words
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestPutID(t *testing.T) {
	if got := PutID([]core.Word{9}, 3); !reflect.DeepEqual(got, []core.Word{9, 0x0088, 3}) {
		t.Fatalf("got %04X", got)
	}
	if got := PutID(nil, 0x10001); !reflect.DeepEqual(got, []core.Word{0x0108, 1, 1}) {
		t.Fatalf("got %04X", got)
	}
}
