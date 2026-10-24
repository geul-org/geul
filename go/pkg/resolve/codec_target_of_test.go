//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestCodecTargetOf — a codec spec document ID gives its target; ordinary documents, 0, a Length and bad hex do not
package resolve

import "testing"

func TestCodecTargetOf(t *testing.T) {
	if w, ok := codecTargetOf("0x0000000000001018"); !ok || w != 0x1018 {
		t.Fatalf("got %04X %v", w, ok)
	}
	for _, s := range []string{"0xffff000000000001", "0x0000000000000000", "0x0000000000001098", "0x00000000000010", "0X0000000000001018"} {
		if _, ok := codecTargetOf(s); ok {
			t.Fatalf("%s: want false", s)
		}
	}
}
