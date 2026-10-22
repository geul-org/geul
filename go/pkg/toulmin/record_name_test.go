//ff:func feature=toulmin type=helper control=sequence
//ff:what TestRecordName — direct codecs and format numbers name; unknown refs are empty
package toulmin

import "testing"

func TestRecordName(t *testing.T) {
	if RecordName(CodecRef{Codec: CodecDoOnce}) != "DO-ONCE" || RecordName(CodecRef{Codec: 0x7C, Format: FormatLabel}) != "LABEL" {
		t.Fatal("name mismatch")
	}
	if RecordName(CodecRef{Codec: 0x66}) != "" || RecordName(CodecRef{Codec: 0x7C, Format: 0xCFFF}) != "" {
		t.Fatal("unknown refs must be empty")
	}
}
