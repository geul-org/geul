//ff:func feature=notation type=encoder control=sequence
//ff:what parseNumbered — #n, Qn or Pn as a LOCAL or Wikidata Reference (LOCAL through the label map)
package notation

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseNumbered reads the prefix (#, Q or P) and a decimal target ID. A LOCAL
// ID goes through ctx.Local when set.
func parseNumbered(tok string, ctx Context) (core.Record, error) {
	n, err := strconv.ParseUint(tok[1:], 10, 64)
	if err != nil {
		return core.Record{}, fmt.Errorf("geul/notation: %q needs a decimal number after %c", tok, tok[0])
	}
	c := map[byte]uint8{'#': core.RefLocal, 'Q': core.RefWikidataQID, 'P': core.RefWikidataPID}[tok[0]]
	if c == core.RefLocal && ctx.Local != nil {
		n = ctx.Local(n)
	}
	return core.Record{Kind: core.KindReference, Codec: c << 3, Payload: core.EncodeUintValue(n).Payload}, nil
}
