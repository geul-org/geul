//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeDecimalBig — SIGNED-DECIMAL 23.45 and UNSIGNED-DECIMAL 1256.1 decode with their quantum; short payloads, other codecs and kinds fail
package core

import "testing"

func TestDecodeDecimalBig(t *testing.T) {
	if e, c, err := DecodeDecimalBig(Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: []Word{0xFFFF, 0xFFFE, 0x0929}}); err != nil || e != -2 || c.Int64() != 2345 {
		t.Fatalf("signed: %d %v %v", e, c, err)
	}
	if e, c, err := DecodeDecimalBig(Record{Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: []Word{0xFFFF, 0xFFFF, 0x3111}}); err != nil || e != -1 || c.Int64() != 12561 {
		t.Fatalf("unsigned: %d %v %v", e, c, err)
	}
	if _, _, err := DecodeDecimalBig(Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: []Word{0, 0}}); err == nil {
		t.Fatal("short: want error")
	}
	if _, _, err := DecodeDecimalBig(Record{Kind: KindReference, Codec: CodecSignedDecimal, Payload: []Word{0, 0, 1}}); err == nil {
		t.Fatal("kind: want error")
	}
	if _, _, err := DecodeDecimalBig(Record{Kind: KindValue, Codec: CodecSignedInteger, Payload: []Word{0, 0, 1}}); err == nil {
		t.Fatal("codec: want error")
	}
}
