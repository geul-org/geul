//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDecimalValue — signed/unsigned layouts and the negative-unsigned error
package core

import "testing"

func TestEncodeDecimalValue(t *testing.T) {
	r, err := EncodeDecimalValue(true, -1, -8)
	if err != nil || r.Codec != CodecSignedDecimal || r.Payload[0] != 0xFFFF || r.Payload[1] != 0xFFFF || r.Payload[2] != 0xFFF8 {
		t.Fatalf("signed: %+v %v", r, err)
	}
	r, err = EncodeDecimalValue(false, 0, 75)
	if err != nil || r.Codec != CodecUnsignedDecimal || len(r.Payload) != 3 || r.Payload[2] != 75 {
		t.Fatalf("unsigned: %+v %v", r, err)
	}
	if _, err := EncodeDecimalValue(false, 0, -1); err == nil {
		t.Fatal("negative unsigned must fail")
	}
}
