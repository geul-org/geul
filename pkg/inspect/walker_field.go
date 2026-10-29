//ff:func feature=inspect type=decoder control=selection
//ff:what walker.field — take one declared field: raw words with their ENUM name, or a record checked against its CODEC targets (a slot marker noted as Marker)
package inspect

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// field takes f at the cursor. A rest field takes every word left but the
// tail fixed fields after it, and adds no part when that is none. A raw field
// is f.Width words noted with the
// ENUM name of that code and, under a CODEC, the Value they are the Payload
// of (rawData); a record field is one record, which must fit f's CODEC
// targets (core.FieldAccepts), noted with its decoded Value (a
// LANGUAGE-SHORT-CODE language named by the registry, namedLanguage), or as a
// Marker when it is a slot marker.
func (w *walker) field(f codecspec.Field) {
	n := len(w.out)
	switch {
	case f.Width == int(core.FieldRest) && w.pos+w.tail == len(w.ws):
		return
	case f.Width == int(core.FieldRest):
		w.words(f.Name, len(w.ws)-w.pos-w.tail)
	case f.Width > 0:
		w.words(f.Name, f.Width)
	default:
		w.rec(f.Name)
	}
	if w.bad || len(w.out) == n {
		return
	}
	p := &w.out[n]
	if f.Width != 0 {
		p.Enum = f.Enums[core.FormatID(p.Words)]
		p.Data = rawData(f.Codecs, p.Words)
		return
	}
	r, _, _ := core.DecodeRecord(p.Words)
	w.bad = !core.FieldAccepts(f.Codecs, r)
	if core.IsSlotMarker(r) {
		p.Data = Marker{Number: r.Codec, Name: core.SlotMarkerName(r.Codec)}
		return
	}
	p.Data = asData(namedLanguage(r, w.langs))
}
