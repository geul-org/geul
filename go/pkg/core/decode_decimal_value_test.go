//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeDecimalValue — signed/unsigned round trips plus short, wrong-codec and int64 overflow errors
package core

import "testing"

func TestDecodeDecimalValue(t *testing.T) {
	r, _ := EncodeDecimalValue(true, -3, -125)
	if e, c, err := DecodeDecimalValue(r); err != nil || e != -3 || c != -125 {
		t.Fatalf("signed: %d %d %v", e, c, err)
	}
	r, _ = EncodeDecimalValue(false, 2, 9)
	if e, c, err := DecodeDecimalValue(r); err != nil || e != 2 || c != 9 {
		t.Fatalf("unsigned: %d %d %v", e, c, err)
	}
	if _, _, err := DecodeDecimalValue(Record{Kind: KindValue, Codec: CodecSignedDecimal, Payload: []Word{0, 0}}); err == nil {
		t.Fatal("short payload must fail")
	}
	if _, _, err := DecodeDecimalValue(Record{Kind: KindValue, Codec: 0x02, Payload: []Word{0, 0, 1}}); err == nil {
		t.Fatal("wrong codec must fail")
	}
	big := Record{Kind: KindValue, Codec: CodecUnsignedDecimal, Payload: []Word{0, 0, 0x8000, 0, 0, 0}}
	if _, _, err := DecodeDecimalValue(big); err == nil {
		t.Fatal("unsigned coefficient above int64 must fail")
	}
}
