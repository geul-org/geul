//ff:func feature=geulmd type=encoder control=selection
//ff:what structuredLine — the typed line of a BODY record and its re-encoding, or an error to fall back
package geulmd

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// structuredLine renders r in typed form and returns the words that form
// encodes to: `VALUE <literal>` for a known standalone Value, or, under the
// TANGEUL profile, the TOULMIN record line. A SOURCE updates the SPAN width in
// w. Anything without a typed form is an error.
func structuredLine(r core.Record, ctx headCtx, w *toulmin.Wire) (string, []core.Word, error) {
	if r.Kind == core.KindValue {
		lit, ok := renderValue(r)
		if !ok {
			return "", nil, errors.New("geulmd: unknown Value")
		}
		enc, _ := core.EncodeRecord(r)
		return "VALUE " + lit, enc, nil
	}
	ref, err := toulmin.RecordRef(r)
	name := toulminName(ref)
	switch {
	case !ctx.tangeul || err != nil || name == "":
		return "", nil, errors.New("geulmd: no structured form")
	case ref.Format == toulmin.FormatSource:
		line, enc, sw, err := sourceForm(r, *w)
		if err == nil {
			w.Span = sw
		}
		return line, enc, err
	default:
		return packetForm(name, ref, r, *w)
	}
}
