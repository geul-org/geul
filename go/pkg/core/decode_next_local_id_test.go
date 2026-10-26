//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeNextLocalID — round trip; another Meta, a non-minimal payload and 0 fail
package core

import "testing"

func TestDecodeNextLocalID(t *testing.T) {
	if n, err := DecodeNextLocalID(EncodeNextLocalID(70000)); err != nil || n != 70000 {
		t.Fatalf("%d %v", n, err)
	}
	for _, r := range []Record{EncodeSpecVersion(2), {Kind: KindMeta, Codec: MetaNextLocalID, Payload: []Word{0, 1}}, EncodeNextLocalID(0)} {
		if _, err := DecodeNextLocalID(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
