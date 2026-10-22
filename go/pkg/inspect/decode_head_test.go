//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecodeHead — ID-WIDTH after IMPORT-DOC still governs it; PROFILE text is returned; malformed widths are ignored
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeHead(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("p-1.0")))
	head := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0, 3, 1, 0, 0, 0, 2}},
		{Kind: core.KindMeta, Codec: core.MetaIDWidth, Payload: []core.Word{0x0400}},
		{Kind: core.KindMeta, Codec: core.MetaIDWidth, Payload: []core.Word{0x0001}},
		{Kind: core.KindMeta, Codec: core.MetaProfile, Payload: vw},
		{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: vw},
	}
	entries, w, profile := decodeHead(head)
	imp, ok := entries[0].Data.(ImportDoc)
	if len(entries) != 5 || !ok || imp.ID != 3 || w != (Widths{1, 1, 2}) || profile != "p-1.0" {
		t.Fatalf("got %+v %+v %q", entries, w, profile)
	}
}
