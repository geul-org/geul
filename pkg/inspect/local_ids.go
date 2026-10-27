//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what LocalIDs — every LOCAL ID a decoded document uses, in order: Edge Ids and the LOCALs records mention
package inspect

// LocalIDs returns, in document order and possibly repeated, each entry's
// Edge Id (when it has one) and the LOCAL IDs it mentions (Mentions): the IDs
// a version uses, which NEXT-LOCAL-ID never gives out again once dropped
// (meta/lineage/NEXT-LOCAL-ID). A LOCAL inside a Value's own payload is not
// looked into.
func LocalIDs(d Document) []uint64 {
	var out []uint64
	for _, e := range append(append([]Entry{}, d.Head...), d.Body...) {
		if e.ID != 0 {
			out = append(out, e.ID)
		}
		out = append(out, Mentions(e)...)
	}
	return out
}
