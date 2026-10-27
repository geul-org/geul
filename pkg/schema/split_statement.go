//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what splitStatement — a BODY record as a statement when it is a TRIPLE or QUALIFIER whose Edge Id and three records decode
package schema

import "github.com/geul-org/gogeul/pkg/core"

// splitStatement cuts r when it is a direct-code TRIPLE or QUALIFIER Edge:
// the Edge Id (edgeID), then three whole records, then the words left. It is
// false for any other record or a payload that does not split.
func splitStatement(r rec) (statement, bool) {
	id, ok := edgeID(r)
	if r.R.Kind != core.KindEdge || (r.R.Codec != core.EdgeTriple && r.R.Codec != core.EdgeQualifier) || !ok {
		return statement{}, false
	}
	_, n, _ := core.DecodeRecord(r.R.Payload)
	ws := r.R.Payload[n:]
	st := statement{R: r, Codec: r.R.Codec, ID: id}
	for i := range st.Parts {
		_, m, err := core.DecodeRecord(ws)
		if err != nil {
			return statement{}, false
		}
		st.Parts[i], ws = ws[:m], ws[m:]
	}
	st.Rest = ws
	return st, true
}
