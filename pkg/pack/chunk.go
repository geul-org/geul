//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what chunk — sorted entries cut into leaf blocks of at most size entries
package pack

// chunk makes leaf blocks of kind from kvs.
func chunk(kvs []kv, kind byte, size int) []block {
	var out []block
	for i, e := range kvs {
		if i%size == 0 {
			out = append(out, block{kind: kind})
		}
		b := &out[len(out)-1]
		b.keys, b.locs = append(b.keys, e.key), append(b.locs, e.loc)
	}
	return out
}
