//ff:func feature=notation type=helper control=iteration dimension=1
//ff:what sampleDoc — test fixture: a document with SPEC-VERSION, LABEL #1 Q42 "Douglas Adams"@en and TRIPLE #2 P31 Q42 Q5 rank 1
package notation

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
)

// sampleDoc encodes a small Wikidata-like document without codec specs.
func sampleDoc() []byte {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	q := func(c uint8, n core.Word) core.Record {
		return core.Record{Kind: core.KindReference, Codec: c << 3, Payload: []core.Word{n}}
	}
	lang := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en"))
	label := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: slices.Concat(enc(core.EncodeLocalRef(1)), enc(q(core.RefWikidataQID, 42)), enc(core.EncodeLanguageText(lang, []byte("Douglas Adams"))))}
	triple := core.Record{Kind: core.KindEdge, Codec: 0x0A, Payload: append(slices.Concat(enc(core.EncodeLocalRef(2)), enc(q(core.RefWikidataPID, 31)), enc(q(core.RefWikidataQID, 42)), enc(q(core.RefWikidataQID, 5))), 1)}
	in := enc(core.EncodeSpecVersion(2))
	in = append(in, 0x180E)
	for _, r := range []core.Record{label, triple} {
		in = append(in, enc(r)...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x180F)))
}
