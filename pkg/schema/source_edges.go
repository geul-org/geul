//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what sourceEdges — E8 over a document: every direct-code SOURCE Edge of the HEAD and BODY checked against the BODY's Edge Ids
package schema

import "github.com/geul-org/gogeul/pkg/core"

// sourceEdges applies sourceFacts to every direct-code SOURCE Edge of the
// HEAD and the BODY, against the BODY's Edge Ids (a SOURCE targets a BODY
// Edge or, from the HEAD, the whole document). The grammar fixes its layout
// (edge/SOURCE), so layer 1 checks it.
func sourceEdges(parts docParts) []Fact {
	ids := edgeIDs(parts.Body)
	var fs []Fact
	for _, r := range append(append([]rec{}, parts.Head...), parts.Body...) {
		if r.R.Kind == core.KindEdge && r.R.Codec == core.EdgeSource && len(r.R.Issuer) == 0 {
			fs = append(fs, sourceFacts(r, ids)...)
		}
	}
	return fs
}
