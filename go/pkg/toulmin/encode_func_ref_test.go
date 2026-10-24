//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeFuncRef — Alias LOCAL REF then VerbCode
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestEncodeFuncRef(t *testing.T) {
	if got := encodeFuncRef(nil, FuncRef{Alias: 2, Verb: 0x7400}); !reflect.DeepEqual(got, []core.Word{0x0088, 2, 0x7400}) {
		t.Fatalf("got %04X", got)
	}
}
