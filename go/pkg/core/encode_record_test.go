//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeRecord — golden TAG/ISSUER layouts and the codec/Issuer validation errors
package core

import "testing"

func TestEncodeRecord(t *testing.T) {
	got, err := EncodeRecord(Record{Kind: KindMeta, Codec: MetaEnd})
	if err != nil || len(got) != 1 || got[0] != 0x180F {
		t.Fatalf("END: got %v err=%v", got, err)
	}
	got, err = EncodeRecord(Record{Kind: KindEdge, Codec: CodecIssuer1Word, Issuer: []Word{0}, Format: 0xC001, Payload: []Word{1}})
	want := []Word{0x11FC, 0x0000, 0xC001, 0x0001}
	if err != nil || len(got) != 4 || got[0] != want[0] || got[2] != want[2] || got[3] != want[3] {
		t.Fatalf("ISSUER edge: got %04X err=%v", got, err)
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: 0x80}); err == nil {
		t.Fatal("codec > 7 bits must fail")
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: CodecIssuer2Word, Issuer: []Word{0}}); err == nil {
		t.Fatal("issuer width mismatch must fail")
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: 0x60, Format: 1}); err == nil {
		t.Fatal("format on a direct codec must fail")
	}
}
