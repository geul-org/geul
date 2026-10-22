//ff:func feature=geulmd type=encoder control=sequence
//ff:what sourceForm — SOURCE as `SOURCE #id span=W "text"` plus its re-encoding and declared span width
package geulmd

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// sourceForm decodes the SOURCE record, renders it with its declared SPAN width
// and quoted text, and re-encodes it so the caller can confirm the line is
// lossless. It also returns the SPAN width for the records that follow.
func sourceForm(r core.Record, w toulmin.Wire) (string, []core.Word, int, error) {
	p, sw, err := toulmin.DecodeSource(r, w.ID)
	if err != nil {
		return "", nil, 0, err
	}
	enc, err := toulmin.EncodeSource(p, toulmin.Wire{ID: w.ID, Span: sw})
	line := fmt.Sprintf("SOURCE #%d span=%d %s", p.TID, sw, strconv.Quote(string(p.Bytes)))
	return line, enc, sw, err
}
