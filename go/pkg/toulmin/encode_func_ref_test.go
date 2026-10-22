//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeFuncRef — Alias REF then VerbCode; an alias past the width fails
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestEncodeFuncRef(t *testing.T) {
	if got, err := encodeFuncRef(nil, FuncRef{Alias: 2, Verb: 0x7400}, testWire); err != nil || !reflect.DeepEqual(got, []core.Word{2, 0x7400}) {
		t.Fatalf("got %04X %v", got, err)
	}
	if _, err := encodeFuncRef(nil, FuncRef{Alias: 0x10000}, testWire); err == nil {
		t.Fatal("alias past width: want error")
	}
}
