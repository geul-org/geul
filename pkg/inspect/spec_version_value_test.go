//ff:func feature=inspect type=decoder control=sequence
//ff:what TestSpecVersionValue — t002 reads as Uint 2; a non-minimal payload is not ok
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecVersionValue(t *testing.T) {
	if v, ok := specVersionValue(core.EncodeSpecVersion(2)); !ok || v.Uint != 2 || v.Codec != core.CodecUnsignedInteger {
		t.Fatalf("got %+v %v", v, ok)
	}
	if _, ok := specVersionValue(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: []core.Word{0, 2}}); ok {
		t.Fatal("non-minimal: want not ok")
	}
}
