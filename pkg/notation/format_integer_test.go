//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatInteger — SIGNED-INTEGER and UNSIGNED-INTEGER past 64 bits by name and digits; a non-minimal payload is false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatInteger(t *testing.T) {
	if s, ok := formatInteger(core.Record{Kind: core.KindValue, Codec: core.CodecSignedInteger, Payload: []core.Word{0x8000}}); !ok || s != "SIGNED-INTEGER(-32768)" {
		t.Fatalf("signed: %s %v", s, ok)
	}
	if s, ok := formatInteger(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: []core.Word{1, 0, 0, 0, 0}}); !ok || s != "UNSIGNED-INTEGER(18446744073709551616)" {
		t.Fatalf("unsigned: %s %v", s, ok)
	}
	if _, ok := formatInteger(core.Record{Kind: core.KindValue, Codec: core.CodecSignedInteger, Payload: []core.Word{0, 1}}); ok {
		t.Fatal("non-minimal: want false")
	}
}
