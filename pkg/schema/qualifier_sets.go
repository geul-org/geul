//ff:func feature=schema type=helper control=iteration dimension=1
//ff:what qualifierSets — for each Edge Id, the set of QUALIFIERs (Property=Value) pointing at it, as one sorted key
package schema

import (
	"slices"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// qualifierSets maps each Target Edge Id to its QUALIFIERs written as
// "Property=Value" in hex, sorted and joined: two Edges with the same key are
// qualified alike (their qualifiers do not tell them apart).
func qualifierSets(sts []statement) map[uint64]string {
	m := map[uint64][]string{}
	for _, s := range sts {
		target, _, err := core.DecodeRecord(s.Parts[1])
		id, lerr := core.DecodeLocalRef(target)
		if s.Codec != core.EdgeQualifier || len(s.Rest) != 0 || err != nil || lerr != nil {
			continue
		}
		m[id] = append(m[id], core.FormatID(s.Parts[0])+"="+core.FormatID(s.Parts[2]))
	}
	out := map[uint64]string{}
	for id, ks := range m {
		slices.Sort(ks)
		out[id] = strings.Join(ks, ";")
	}
	return out
}
