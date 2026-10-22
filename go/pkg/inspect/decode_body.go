//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what decodeBody — BODY entries in order, the SOURCE setting the SPAN width for the records after it
package inspect

import (
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// decodeBody interprets each BODY record in order. idWidth is the declared
// Edge Id width; the SPAN width starts at 1 word until a SOURCE declares it.
// tg turns on TOULMIN decoding.
func decodeBody(body []core.Record, idWidth int, tg bool) []Entry {
	w := toulmin.Wire{ID: idWidth, Span: 1}
	out := make([]Entry, 0, len(body))
	for _, r := range body {
		out = append(out, bodyEntry(r, tg, &w))
	}
	return out
}
