//ff:func feature=delta type=helper control=iteration dimension=1
//ff:what removedIDs — the base Edge Ids the next version no longer has, in base order
package delta

import "slices"

// removedIDs returns the IDs of base that next lacks, in base order.
func removedIDs(base, next []uint64) []uint64 {
	var out []uint64
	for _, id := range base {
		if !slices.Contains(next, id) {
			out = append(out, id)
		}
	}
	return out
}
