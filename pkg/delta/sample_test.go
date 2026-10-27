//ff:func feature=delta type=helper control=sequence
//ff:what bindOf — test fixture: a BIND Edge #id naming LOCAL subject with text
package delta

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// bindOf is BIND(#id, LOCAL subject, text) for tests.
func bindOf(id, subject uint64, text string) core.Record {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := slices.Concat(enc(core.EncodeLocalRef(id)), enc(core.EncodeLocalRef(subject)), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(text))))
	return core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p}
}
