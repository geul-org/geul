//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeIntegerBig — SIGNED-INTEGER -1 and UNSIGNED-INTEGER 0xFFFF decode; other codecs and non-minimal payloads fail
package core

import "testing"

func TestDecodeIntegerBig(t *testing.T) {
	if n, err := DecodeIntegerBig(Record{Kind: KindValue, Codec: CodecSignedInteger, Payload: []Word{0xFFFF}}); err != nil || n.Int64() != -1 {
		t.Fatalf("signed: %v %v", n, err)
	}
	if n, err := DecodeIntegerBig(Record{Kind: KindValue, Codec: CodecUnsignedInteger, Payload: []Word{0xFFFF}}); err != nil || n.Int64() != 0xFFFF {
		t.Fatalf("unsigned: %v %v", n, err)
	}
	if _, err := DecodeIntegerBig(Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: []Word{0, 0, 1}}); err == nil {
		t.Fatal("decimal: want error")
	}
	if _, err := DecodeIntegerBig(Record{Kind: KindReference, Codec: CodecSignedInteger, Payload: []Word{1}}); err == nil {
		t.Fatal("reference: want error")
	}
	if _, err := DecodeIntegerBig(Record{Kind: KindValue, Codec: CodecSignedInteger, Payload: []Word{0, 1}}); err == nil {
		t.Fatal("non-minimal: want error")
	}
}
