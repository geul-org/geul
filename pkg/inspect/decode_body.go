//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what decodeBody — BODY entries in order, the SOURCE setting the SPAN width for the records after it
package inspect

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// decodeBody interprets each BODY record in order. The SPAN width starts at 1
// word until a SOURCE declares it. tg turns on TOULMIN decoding.
func decodeBody(body []core.Record, tg bool) []Entry {
	w := toulmin.Wire{Span: 1}
	out := make([]Entry, 0, len(body))
	for _, r := range body {
		out = append(out, bodyEntry(r, tg, &w))
	}
	return out
}
