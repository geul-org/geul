//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeRecord — round trips across Length forms plus prefix, truncation and ISSUER errors
package core

import "testing"

func TestDecodeRecord(t *testing.T) {
	for _, n := range []int{0, 13, 14, 0xFFFD, 0xFFFE} {
		in := Record{Kind: KindValue, Codec: 0x40, Payload: make([]Word, n)}
		ws, err := EncodeRecord(in)
		if err != nil {
			t.Fatal(err)
		}
		got, used, err := DecodeRecord(ws)
		if err != nil || used != len(ws) || len(got.Payload) != n || got.Codec != 0x40 || got.Kind != KindValue {
			t.Fatalf("n=%d: used=%d err=%v", n, used, err)
		}
	}
	iss, _ := EncodeRecord(Record{Kind: KindEdge, Codec: CodecIssuer1Word, Issuer: []Word{7}, Format: 0xC005, Payload: []Word{9}})
	got, _, err := DecodeRecord(iss)
	if err != nil || got.Issuer[0] != 7 || got.Format != 0xC005 || got.Payload[0] != 9 {
		t.Fatalf("issuer decode: %+v %v", got, err)
	}
	bad := [][]Word{
		{},
		{0x2000},
		{0x11FC},
		{0x1080},
		{0x107C},
		{0x007C | 2<<11 | 1<<7},
	}
	for _, b := range bad {
		if _, _, err := DecodeRecord(b); err == nil {
			t.Fatalf("expected error for %04X", b)
		}
	}
	if _, _, err := DecodeRecord([]Word{0x1000 | 0xE<<7, 3, 0}); err == nil {
		t.Fatal("total smaller than header must fail")
	}
}
