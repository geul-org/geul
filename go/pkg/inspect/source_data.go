//ff:func feature=inspect type=decoder control=sequence
//ff:what sourceData — the SOURCE as Source, its re-encoding, and the SPAN width it sets for later records
package inspect

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// sourceData decodes the SOURCE record and re-encodes it at its declared SPAN
// width. On success the width becomes w.Span for the records after it.
func sourceData(r core.Record, w *toulmin.Wire) (any, []core.Word, error) {
	p, sw, err := toulmin.DecodeSource(r, w.ID)
	if err != nil {
		return nil, nil, err
	}
	enc, err := toulmin.EncodeSource(p, toulmin.Wire{ID: w.ID, Span: sw})
	if err == nil {
		w.Span = sw
	}
	return Source{TextPacket: p, SpanWidth: sw}, enc, err
}
