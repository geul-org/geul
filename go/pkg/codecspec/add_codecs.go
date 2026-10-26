//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what addCodecs — attach each CODEC's target to its field; a raw field keeps only its first
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// addCodecs puts each CODEC target of body on its field in document order
// (edge/CODEC). A raw field takes one CODEC; a later one is skipped (schema E5
// reports it), as is a CODEC of an unknown field.
func addCodecs(body []core.Record, fields map[uint64]*Field) {
	for _, r := range body {
		c, err := core.DecodeCodecDecl(r)
		if err != nil {
			continue
		}
		if f := fields[c.Field]; f != nil && (f.Width == 0 || len(f.Codecs) == 0) {
			f.Codecs = append(f.Codecs, c.Target)
		}
	}
}
