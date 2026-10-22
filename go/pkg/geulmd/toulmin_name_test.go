//ff:func feature=geulmd type=helper control=sequence
//ff:what TestToulminName — direct codecs and format numbers name; unknown refs are empty
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestToulminName(t *testing.T) {
	if toulminName(toulmin.CodecRef{Codec: toulmin.CodecDoOnce}) != "DO-ONCE" || toulminName(toulmin.CodecRef{Codec: 0x7C, Format: toulmin.FormatLabel}) != "LABEL" {
		t.Fatal("name mismatch")
	}
	if toulminName(toulmin.CodecRef{Codec: 0x66}) != "" || toulminName(toulmin.CodecRef{Codec: 0x7C, Format: 0xCFFF}) != "" {
		t.Fatal("unknown refs must be empty")
	}
}
