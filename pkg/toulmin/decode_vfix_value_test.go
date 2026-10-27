//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeVFixValue — round trips plus wrong codec, short payload and out-of-range errors
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeVFixValue(t *testing.T) {
	for _, v := range []VFix{0, 10000, -10000, 8000, 1} {
		if got, err := DecodeVFixValue(EncodeVFixValue(v)); err != nil || got != v {
			t.Fatalf("%d: got %d %v", v, got, err)
		}
	}
	if _, err := DecodeVFixValue(core.EncodeUintValue(1)); err == nil {
		t.Fatal("wrong codec must fail")
	}
	if _, err := DecodeVFixValue(core.Record{Kind: core.KindValue, Codec: core.CodecSignedDecimal}); err == nil {
		t.Fatal("short payload must fail")
	}
	r, _ := core.EncodeDecimalValue(true, 0, 2)
	if _, err := DecodeVFixValue(r); err == nil {
		t.Fatal("2.0 is outside the verdict range")
	}
}
