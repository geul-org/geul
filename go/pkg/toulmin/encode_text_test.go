//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeText — TEXT golden bytes and the ID-width error
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

// TestEncodeText pins the TEXT record of "abc" (ISSUER-1WORD, GEUL Issuer, format
// 0xC001, Edge Id 2, UTF8-TEXT-ODD Value) and checks an Edge Id over the width.
func TestEncodeText(t *testing.T) {
	got, err := EncodeText(TextPacket{TID: 2, Bytes: []byte("abc")}, testWire)
	if want := []core.Word{0x137C, 0, 0xC001, 2, 0x0941, 0x6162, 0x6300}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X %v want %04X", got, err, want)
	}
	if _, err := EncodeText(TextPacket{TID: 0x10000}, testWire); err == nil {
		t.Fatal("ID over width: want error")
	}
}
