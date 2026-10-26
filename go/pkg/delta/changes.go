//ff:func feature=delta type=decoder control=sequence
//ff:what Changes — the differences between two versions of a document: HEAD records added or removed, Edges matched by Edge Id
package delta

// Changes compares base and next, two snapshots of one document whose BODY
// is all Edges with LOCAL Edge Ids. HEAD records are compared as a list of
// whole records (headChanges). Edges are matched by Edge Id, which stays with
// a statement across versions (OVERVIEW 개정): in base order, removed ('-')
// and changed ('~'); then, in next order, added ('+').
func Changes(base, next []byte) ([]Change, error) {
	b, err := splitSnapshot(base)
	if err != nil {
		return nil, err
	}
	n, err := splitSnapshot(next)
	if err != nil {
		return nil, err
	}
	out := headChanges(b.head, n.head)
	return append(out, bodyChanges(b, n)...), nil
}
