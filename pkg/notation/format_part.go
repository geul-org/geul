//ff:func feature=notation type=decoder control=selection
//ff:what formatPart — one field as a token: an embedded record by formatRecord, a raw field by its ENUM name or raw words
package notation

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// formatPart writes one payload part.
func formatPart(p inspect.Part) string {
	switch {
	case p.Record:
		r, _, _ := core.DecodeRecord(p.Words)
		return formatRecord(r)
	case p.Enum != "":
		return p.Enum
	default:
		return hexToken(p.Words)
	}
}
