//ff:func feature=inspect type=decoder control=selection
//ff:what walker.field — take one declared field: raw words with their ENUM name, or a record checked against its expected TAG
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// field takes f at the cursor. A rest field takes every word left but the
// tail fixed fields after it, and adds no part when that is none. A raw field
// is f.Width words noted with the
// ENUM name of that code; a record field is one record, which must have the
// expected Kind·Codec when f declares one, noted with its decoded Value.
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
	if f.Width > 0 {
		p.Enum = f.Enums[core.FormatID(p.Words)]
		return
	}
	r, _, _ := core.DecodeRecord(p.Words)
	w.bad = f.HasExpect && core.TargetWord(r.Kind, r.Codec) != f.Expect
	p.Data = asData(parseValue(r))
}
