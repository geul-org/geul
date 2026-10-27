//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeUintValue — round trip and wrong-codec error
package core

import "testing"

func TestDecodeUintValue(t *testing.T) {
	if n, err := DecodeUintValue(EncodeUintValue(1 << 40)); err != nil || n != 1<<40 {
		t.Fatalf("got %d %v", n, err)
	}
	if _, err := DecodeUintValue(Record{Kind: KindValue, Codec: 0x01}); err == nil {
		t.Fatal("wrong codec must fail")
	}
}
