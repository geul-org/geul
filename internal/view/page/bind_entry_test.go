//ff:func feature=page type=builder control=sequence
//ff:what bindEntry — test fixture: the Entry of a BIND naming LOCAL subj, laid out by a BIND codec spec
package page

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// bindEntry builds the Entry of BIND #id naming LOCAL subj with BIND's codec
// spec (subject, name), as DecodeWith would.
func bindEntry(id, subj uint64, name string) inspect.Entry {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := append(enc(core.EncodeLocalRef(id)), enc(core.EncodeLocalRef(subj))...)
	p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(name)))...)
	e := inspectEntry(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p}, []codecspec.Field{{Name: "subject"}, {Name: "name"}})
	e.ID = id
	return e
}
