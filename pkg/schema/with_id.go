//ff:func feature=schema type=helper control=sequence
//ff:what withID — a set of IDs with one more, made when nil
package schema

// withID adds id to set, making the set when it is nil, and returns it.
func withID(set map[uint64]bool, id uint64) map[uint64]bool {
	if set == nil {
		set = map[uint64]bool{}
	}
	set[id] = true
	return set
}
