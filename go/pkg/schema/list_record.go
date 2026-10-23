//ff:func feature=schema type=rule control=selection
//ff:what listRecord — P1, P2 and the record's own rules for one version list BODY record
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// listRecord checks the i-th BODY record of a version list: its Edge Id is i+1
// (P2), and it is a BIND (P5) or a PUBLICATION/DEPRECATION/TOMBSTONE of the
// listed document (P3, P4); anything else is P1.
func listRecord(st *listState, i int, r rec) []Fact {
	var fs []Fact
	if id, ok := edgeID(r); ok && id != uint64(i+1) {
		fs = append(fs, factAt("P2", Error, r, "Edge Id", fmt.Sprintf("%d", i+1), fmt.Sprintf("%d", id)))
	}
	switch {
	case r.R.Kind == core.KindEdge && r.R.Codec == core.EdgeBind:
		return append(fs, applyBind(st, r)...)
	default:
		p, err := core.DecodePublication(r.R)
		if err != nil || core.FormatID(p.Doc) != st.doc {
			return append(fs, factAt("P1", Error, r, "", "이 문서("+st.doc+")의 PUBLICATION·DEPRECATION·TOMBSTONE·BIND", fmt.Sprintf("kind %d codec 0x%02X", r.R.Kind, r.R.Codec)))
		}
		return append(fs, applyPublication(st.docs, r, p)...)
	}
}
