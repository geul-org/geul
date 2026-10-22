//ff:func feature=toulmin type=decoder control=sequence
//ff:what TestDecodeSpan — round trip at width 2 plus truncated offset and length errors
package toulmin

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeSpan(t *testing.T) {
	w, _ := EncodeSpan(Span{Off: 70000, Len: 5}, 2)
	if s, err := DecodeSpan(w, 2); err != nil || s.Off != 70000 || s.Len != 5 {
		t.Fatalf("got %+v %v", s, err)
	}
	if _, err := DecodeSpan(nil, 1); err == nil {
		t.Fatal("missing offset must fail")
	}
	if _, err := DecodeSpan([]core.Word{1}, 1); err == nil {
		t.Fatal("missing length must fail")
	}
}
