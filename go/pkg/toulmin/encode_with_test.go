//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeWith — count then REFs; a failing item before the last and at the end both error
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestEncodeWith(t *testing.T) {
	if got, err := encodeWith(nil, []TID{5, 6}, testWire); err != nil || !reflect.DeepEqual(got, []core.Word{2, 5, 6}) {
		t.Fatalf("got %v %v", got, err)
	}
	if _, err := encodeWith(nil, []TID{0x10000, 1}, testWire); err == nil {
		t.Fatal("bad first item: want error")
	}
	if _, err := encodeWith(nil, []TID{0x10000}, testWire); err == nil {
		t.Fatal("bad last item: want error")
	}
}
