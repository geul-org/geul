//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckBytes — even and odd codecs, empty odd payload, non-zero pad
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckBytes(t *testing.T) {
	v := func(c uint8, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindValue, Codec: c, Payload: p})
	}
	if checkBytes(recOf(core.EncodeBytesValue(0x40, []byte("abc")))) != nil || checkBytes(v(0x40)) != nil {
		t.Fatal("ok cases")
	}
	if rulesOf(checkBytes(v(0x41))) != "V4" || rulesOf(checkBytes(v(0x41, 0x6162))) != "V4" {
		t.Fatal("pad cases")
	}
	if rulesOf(checkBytes(v(0x45, 0x2000))) != "V5" {
		t.Fatal("odd URI of a space")
	}
}
