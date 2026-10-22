//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeBytesValue — round trip plus wrong codec and bad pad errors
package core

import "testing"

func TestDecodeBytesValue(t *testing.T) {
	for _, s := range []string{"", "a", "ab", "한글"} {
		b, err := DecodeBytesValue(EncodeBytesValue(CodecUTF8TextEven, []byte(s)), CodecUTF8TextEven)
		if err != nil || string(b) != s {
			t.Fatalf("%q: got %q err=%v", s, b, err)
		}
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: 0x46}, CodecUTF8TextEven); err == nil {
		t.Fatal("wrong codec must fail")
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: 0x41, Payload: []Word{0x6162}}, CodecUTF8TextEven); err == nil {
		t.Fatal("non-zero pad must fail")
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: 0x41}, CodecUTF8TextEven); err == nil {
		t.Fatal("odd without payload must fail")
	}
}
