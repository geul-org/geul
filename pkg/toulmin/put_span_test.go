//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestPutSpan — appends offset and length; overflow fails
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPutSpan(t *testing.T) {
	if got, err := PutSpan(nil, Span{Off: 1, Len: 2}, 1); err != nil || !reflect.DeepEqual(got, []core.Word{1, 2}) {
		t.Fatalf("got %v %v", got, err)
	}
	if _, err := PutSpan(nil, Span{Off: 0x10000}, 1); err == nil {
		t.Fatal("overflow: want error")
	}
}
