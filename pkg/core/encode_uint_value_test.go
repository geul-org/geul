//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeUintValue — UNSIGNED-INTEGER codec and minimal payload
package core

import "testing"

func TestEncodeUintValue(t *testing.T) {
	r := EncodeUintValue(70000)
	if r.Kind != KindValue || r.Codec != CodecUnsignedInteger || len(r.Payload) != 2 {
		t.Fatalf("%+v", r)
	}
}
