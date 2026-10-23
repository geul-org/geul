//ff:func feature=resolve type=encoder control=iteration dimension=1
//ff:what sampleDoc — test fixture: a named document under Issuer 0x0000 with IMPORT-DOC-1WORD records and one text Value
package resolve

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// sampleDoc builds document 0xffff00000000<id> of Issuer 0x0000 whose BODY is
// text. Each import is {IMPORT ID, imported document id, version ID}.
func sampleDoc(id core.Word, text string, imports ...[3]uint32) []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := enc(core.Record{Kind: core.KindMeta, Codec: core.MetaDocName, Payload: []core.Word{0, 0xffff, 0, 0, id}})
	for _, m := range imports {
		p := append(core.EncodeUintValue(uint64(m[0])).Payload, 0, 0xffff, 0, 0, core.Word(m[1]), core.Word(m[2]>>16), core.Word(m[2]))
		in = append(in, enc(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc1Word, Payload: p})...)
	}
	in = append(append(in, 0x180E), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text)))...)
	return core.PutWords(core.WrapDocument(append(in, 0x180F)))
}
