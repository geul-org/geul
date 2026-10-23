//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what addEnums — attach each ENUM's name to its raw field when the code width fits
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// addEnums puts each ENUM of body on its field: the field must be a raw field
// and the code exactly its width (edge/ENUM). Others are skipped.
func addEnums(body []core.Record, fields map[uint64]*Field, names map[uint64]string) {
	for _, r := range body {
		e, err := core.DecodeEnumDecl(r)
		if err != nil {
			continue
		}
		if f := fields[e.Field]; f != nil && f.Width > 0 && len(e.Code) == f.Width {
			f.Enums[core.FormatID(e.Code)] = names[e.EdgeID]
		}
	}
}
