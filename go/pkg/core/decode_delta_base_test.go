//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDeltaBase — round trip of a 2-word version; another Meta, a wrong width and version 0 fail
package core

import "testing"

func TestDecodeDeltaBase(t *testing.T) {
	if v, err := DecodeDeltaBase(EncodeDeltaBase(0x10002)); err != nil || v != 0x10002 {
		t.Fatalf("%d %v", v, err)
	}
	for _, r := range []Record{EncodeSpecVersion(2), {Kind: KindMeta, Codec: MetaDeltaBase, Payload: []Word{3}}, EncodeDeltaBase(0)} {
		if _, err := DecodeDeltaBase(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
