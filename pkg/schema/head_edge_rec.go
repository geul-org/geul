//ff:func feature=schema type=rule control=sequence
//ff:what headEdgeRec — whether a record is a HEAD Edge (SOURCE on the whole document or a QUALIFIER on one), noting a SOURCE's Edge Id
package schema

import "github.com/geul-org/gogeul/pkg/core"

// headEdgeRec reports whether r is a HEAD Edge given the HEAD SOURCEs so far
// (core.HeadEdge) and, when r is such a SOURCE, adds its Edge Id to
// st.sources.
func headEdgeRec(st *placeState, r rec) bool {
	id, ok := core.HeadEdge(r.R, st.sources)
	if ok && r.R.Codec == core.EdgeSource {
		st.sources = withID(st.sources, id)
	}
	return ok
}
