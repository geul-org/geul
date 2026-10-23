//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeSpecVersion — round trip; another Meta, an empty and a non-minimal payload fail
package core

import "testing"

func TestDecodeSpecVersion(t *testing.T) {
	if v, err := DecodeSpecVersion(EncodeSpecVersion(2)); err != nil || v != 2 {
		t.Fatalf("got %d %v", v, err)
	}
	for _, r := range []Record{
		{Kind: KindMeta, Codec: MetaProfile, Payload: []Word{2}},
		{Kind: KindMeta, Codec: MetaSpecVersion},
		{Kind: KindMeta, Codec: MetaSpecVersion, Payload: []Word{0, 2}},
	} {
		if _, err := DecodeSpecVersion(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
