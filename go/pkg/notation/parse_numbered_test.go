//ff:func feature=notation type=encoder control=sequence
//ff:what TestParseNumbered — #n through the label map, Qn and Pn; a non-number fails
package notation

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestParseNumbered(t *testing.T) {
	ctx := Context{Local: func(n uint64) uint64 { return n + 10 }}
	if r, _ := parseNumbered("#2", ctx); r.Payload[0] != 12 || r.Codec != core.CodecLocalIdentified {
		t.Fatalf("%+v", r)
	}
	if r, _ := parseNumbered("P31", ctx); r.Payload[0] != 31 || r.Codec != core.RefWikidataPID<<3 {
		t.Fatalf("%+v", r)
	}
	if _, err := parseNumbered("#x", ctx); err == nil {
		t.Fatal("want error")
	}
}
