//ff:func feature=codecspec type=encoder control=sequence
//ff:what nameOf — test fixture: a BIND naming a LOCAL Subject
package codecspec

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// nameOf builds a BIND naming the LOCAL Subject subj.
func nameOf(id, subj uint64, name string) core.Record {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := slices.Concat(enc(core.EncodeLocalRef(id)), enc(core.EncodeLocalRef(subj)), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(name))))
	return core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p}
}
