//ff:func feature=inspect type=decoder control=selection
//ff:what bodyEntry — one BODY entry: a known Value, or under TOULMIN a lossless TOULMIN record; bare otherwise
package inspect

import (
	"slices"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// bodyEntry interprets a BODY record. A standalone Value of a known codec
// carries its Value; a version list record (PUBLICATION, DEPRECATION,
// TOMBSTONE) outside TANGEUL its core.Publication and Edge Id. Under the TANGEUL profile (tg) an Edge record takes its
// TOULMIN name and, when it decodes and re-encodes to the same words, its
// packet, Edge Id and references. Everything else stays bare.
func bodyEntry(r core.Record, tg bool, w *toulmin.Wire) Entry {
	e := newEntry(r)
	ref, err := toulmin.RecordRef(r)
	switch {
	case r.Kind == core.KindValue:
		e.Data = asData(parseValue(r))
		return e
	case !tg && r.Kind == core.KindEdge && r.Codec >= core.EdgePublication && r.Codec <= core.EdgeTombstone:
		if p, err := core.DecodePublication(r); err == nil {
			e.Data, e.ID = p, p.EdgeID
		}
		return e
	case !tg || err != nil:
		return e
	}
	if n := toulmin.RecordName(ref); n != "" {
		e.Name = n
	}
	data, enc, err := toulminData(r, ref, w)
	if err == nil && slices.Equal(enc, e.Words) {
		e.Data, e.ID, e.Refs = data, idOf(data), refsOf(data)
	}
	return e
}
