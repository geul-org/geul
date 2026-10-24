//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what declaredFields — the fields of a DECLARE: FIELDs naming it, in document order
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// declaredFields lists the FIELDs of body whose DECLARE is decl, in document
// order (edge/FIELD).
func declaredFields(decl uint64, body []core.Record, fields map[uint64]*Field) []Field {
	var out []Field
	for _, r := range body {
		f, err := core.DecodeFieldDecl(r)
		if err != nil || f.Declare != decl || fields[f.EdgeID] == nil {
			continue
		}
		out = append(out, *fields[f.EdgeID])
	}
	return out
}
