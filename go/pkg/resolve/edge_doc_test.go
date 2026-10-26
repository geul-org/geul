//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what edgeDoc — test fixture: document 0xffff00000000<id> of Issuer 0x0000 whose BODY is BIND #n = text, numbered from 1
package resolve

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
)

// edgeDoc builds a diffable document: DOC-ID, then BIND(#i, LOCAL 100,
// texts[i-1]) for each text.
func edgeDoc(id core.Word, texts ...string) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	doc, _ := core.EncodeDocID([]core.Word{0}, []core.Word{0xffff, 0, 0, id})
	in := append(enc(doc), 0x180E)
	for i, s := range texts {
		p := slices.Concat(enc(core.EncodeLocalRef(uint64(i+1))), enc(core.EncodeLocalRef(100)), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s))))
		in = append(in, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x180F)))
}
