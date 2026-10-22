//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestPutID — appends a width-word field; overflow fails
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestPutID(t *testing.T) {
	if got, err := PutID([]core.Word{9}, 0x10001, 2); err != nil || !reflect.DeepEqual(got, []core.Word{9, 1, 1}) {
		t.Fatalf("got %v %v", got, err)
	}
	if _, err := PutID(nil, 0x10000, 1); err == nil {
		t.Fatal("overflow: want error")
	}
}
