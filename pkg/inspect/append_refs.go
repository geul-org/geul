//ff:func feature=inspect type=helper control=iteration dimension=1
//ff:what appendRefs — append the non-zero Edge Ids (0 is "none")
package inspect

import "github.com/geul-org/gogeul/pkg/toulmin"

// appendRefs appends each non-zero id to out; 0 is the REF "none".
func appendRefs(out []uint64, ids ...toulmin.TID) []uint64 {
	for _, id := range ids {
		if id != 0 {
			out = append(out, uint64(id))
		}
	}
	return out
}
