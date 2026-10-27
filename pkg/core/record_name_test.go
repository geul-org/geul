//ff:func feature=core type=helper control=sequence
//ff:what TestRecordName — a known Value ISSUER format by its format name, an unknown one and a direct codec by codec name
package core

import "testing"

func TestRecordName(t *testing.T) {
	if got := RecordName(Record{Kind: KindValue, Codec: CodecIssuer1Word, Issuer: []Word{0x0001}, Format: 0x0700}); got != "COMPLEX64" {
		t.Fatalf("complex64: %q", got)
	}
	if got := RecordName(Record{Kind: KindValue, Codec: CodecIssuer1Word, Issuer: []Word{0x0009}, Format: 0x0700}); got != "ISSUER-1WORD" {
		t.Fatalf("unknown format: %q", got)
	}
	if got := RecordName(Record{Kind: KindValue, Codec: CodecFraction}); got != "FRACTION" {
		t.Fatalf("fraction: %q", got)
	}
}
