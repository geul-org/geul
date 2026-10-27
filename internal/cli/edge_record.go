//ff:func feature=cli type=helper control=sequence
//ff:what edgeRecord — the Edge record of a parsed line under a given Edge Id
package cli

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

// edgeRecord puts LOCAL id before e's payload.
func edgeRecord(id uint64, e notation.Edge) core.Record {
	lead, _ := core.EncodeRecord(core.EncodeLocalRef(id))
	return core.Record{Kind: core.KindEdge, Codec: e.Codec, Payload: append(lead, e.Rest...)}
}
