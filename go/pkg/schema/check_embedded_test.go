//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckEmbedded — no Value, wrong codec, a Value that breaks V rules, ok
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
	"testing"
)

func TestCheckEmbedded(t *testing.T) {
	m := func(p []core.Word) rec { return recOf(core.Record{Kind: core.KindMeta, Codec: 0x10, Payload: p}) }
	if fs := checkEmbedded(m(nil), "M4", "x", isUTF8Text); rulesOf(fs) != "M4" {
		t.Fatalf("none: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.EncodeUintValue(1))), "M4", "x", isUTF8Text); rulesOf(fs) != "M4" {
		t.Fatalf("codec: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.Record{Kind: core.KindValue, Codec: 0x41})), "M4", "x", isUTF8Text); rulesOf(fs) != "V4" {
		t.Fatalf("value rule: %+v", fs)
	}
	if fs := checkEmbedded(m(enc(core.EncodeBytesValue(0x40, []byte("geul-t002")))), "M4", "x", isUTF8Text); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
}
