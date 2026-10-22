//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeSpan — offset then length at the declared width, and overflow errors per field
package toulmin

import "testing"

func TestEncodeSpan(t *testing.T) {
	w, err := EncodeSpan(Span{Off: 3, Len: 4}, 1)
	if err != nil || len(w) != 2 || w[0] != 3 || w[1] != 4 {
		t.Fatalf("got %v %v", w, err)
	}
	if _, err := EncodeSpan(Span{Off: 0x10000}, 1); err == nil {
		t.Fatal("offset overflow must fail")
	}
	if _, err := EncodeSpan(Span{Len: 0x10000}, 1); err == nil {
		t.Fatal("length overflow must fail")
	}
}
