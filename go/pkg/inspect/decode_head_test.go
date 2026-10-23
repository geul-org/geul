//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecodeHead — every Meta becomes an entry in order and the PROFILE text is returned
package inspect

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/core"
)

func TestDecodeHead(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("p-1.0")))
	head := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaImportDoc1Word, Payload: []core.Word{3, 1, 0, 0, 0, 2, 0, 0}},
		{Kind: core.KindMeta, Codec: core.MetaProfile, Payload: vw},
		{Kind: core.KindMeta, Codec: core.MetaSpecVersion, Payload: vw},
	}
	entries, profile := decodeHead(head)
	imp, ok := entries[0].Data.(ImportDoc)
	if len(entries) != 3 || !ok || imp.ID != 3 || profile != "p-1.0" {
		t.Fatalf("got %+v %q", entries, profile)
	}
}
