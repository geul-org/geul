//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what MaxLocalID — the largest LOCAL ID a decoded document uses: Edge Ids and LOCAL References in its record fields
package inspect

// MaxLocalID returns the largest LOCAL ID d uses (0 when none): each
// entry's Edge Id and every LOCAL Reference among the record parts its
// layout gives (partsLocal). It is what an auto increment must stay above
// (meta/lineage/NEXT-LOCAL-ID). A LOCAL held inside a Value's own payload
// (a globe or unit Reference) is not looked into.
func MaxLocalID(d Document) uint64 {
	var max uint64
	for _, e := range append(append([]Entry{}, d.Head...), d.Body...) {
		max = maxOf(max, e.ID, partsLocal(Layout(e)))
	}
	return max
}
