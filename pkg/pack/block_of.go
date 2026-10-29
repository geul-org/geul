//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what blockOf — a directory block of kind over keys and locations
package pack

// blockOf makes a block of kind whose entries are kvs in order.
func blockOf(kind byte, kvs []streamKV) block {
	b := block{kind: kind, keys: make([]string, len(kvs)), locs: make([]Loc, len(kvs))}
	for i, e := range kvs {
		b.keys[i], b.locs[i] = e.Key, e.Loc
	}
	return b
}
