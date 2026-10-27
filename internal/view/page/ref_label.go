//ff:func feature=page type=helper control=iteration dimension=1
//ff:what refLabel — a node record's payload as its fields in hex, space-separated
package page

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// refLabel shows r's payload field by field: a Reference's named fields
// (refPayloadLines) each in `0x` hex, a lone TAG as nothing, any other
// payload as one hex run.
func refLabel(r core.Record) string {
	if r.Kind != core.KindReference {
		if len(r.Payload) == 0 {
			return ""
		}
		return core.FormatID(r.Payload)
	}
	var fields [][]core.Word
	for i, l := range refPayloadLines(r) {
		if l.Name != "" || len(fields) == 0 {
			fields = append(fields, nil)
		}
		fields[len(fields)-1] = append(fields[len(fields)-1], r.Payload[i])
	}
	parts := make([]string, len(fields))
	for i, f := range fields {
		parts[i] = core.FormatID(f)
	}
	return strings.Join(parts, " ")
}
