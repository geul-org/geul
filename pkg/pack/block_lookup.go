//ff:func feature=pack type=resolver control=sequence
//ff:what block.lookup — the location of key in a leaf, false when absent
package pack

import "sort"

// lookup finds key in a leaf block.
func (b block) lookup(key string) (Loc, bool) {
	i := sort.SearchStrings(b.keys, key)
	if i < len(b.keys) && b.keys[i] == key {
		return b.locs[i], true
	}
	return Loc{}, false
}
