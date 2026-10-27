//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerFormatName — ADHOC 0x0702 is RATIONAL-I64, IETF 0xC001 IPV4-ADDRESS; an Edge, a 2-WORD Issuer, a direct codec and an unassigned format are ""
package core

import "testing"

func TestIssuerFormatName(t *testing.T) {
	if got := IssuerFormatName(Record{Kind: KindValue, Codec: CodecIssuer1Word, Issuer: []Word{0x0001}, Format: 0x0702}); got != "RATIONAL-I64" {
		t.Fatalf("rational: %q", got)
	}
	if got := IssuerFormatName(Record{Kind: KindValue, Codec: CodecIssuer1Word, Issuer: []Word{0x0020}, Format: 0xC001}); got != "IPV4-ADDRESS" {
		t.Fatalf("ipv4: %q", got)
	}
	for _, r := range []Record{
		{Kind: KindEdge, Codec: CodecIssuer1Word, Issuer: []Word{0}, Format: 0xC000},
		{Kind: KindValue, Codec: CodecIssuer2Word, Issuer: []Word{0, 1}, Format: 0x0702},
		{Kind: KindValue, Codec: CodecFraction},
		{Kind: KindValue, Codec: CodecIssuer1Word, Issuer: []Word{0x0001}, Format: 0x0705},
	} {
		if got := IssuerFormatName(r); got != "" {
			t.Fatalf("%+v: %q", r, got)
		}
	}
}
