//ff:func feature=delta type=helper control=sequence
//ff:what refNodeOf — test fixture: a REFERENCE-NODE Edge #id binding LOCAL subject to a URI
package delta

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// refNodeOf is REFERENCE-NODE(#id, LOCAL subject, URI uri) for tests.
func refNodeOf(id, subject uint64, uri string) core.Record {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := slices.Concat(enc(core.EncodeLocalRef(id)), enc(core.EncodeLocalRef(subject)), enc(core.EncodeBytesValue(0x44, []byte(uri))))
	return core.Record{Kind: core.KindEdge, Codec: edgeReferenceNode, Payload: p}
}
