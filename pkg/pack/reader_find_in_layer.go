//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.findInLayer — descend one layer from its root to the leaf entry of key (§5.1), false when absent
package pack

import "context"

// findInLayer returns key's leaf location in l (possibly a deletion marker)
// and whether the layer has the key at all.
func (r *Reader) findInLayer(ctx context.Context, l layer, key string) (Loc, bool, error) {
	b := l.blk
	for depth := 0; depth < maxDepth; depth++ {
		if isLeaf(b.kind) {
			loc, ok := b.lookup(key)
			return loc, ok, nil
		}
		i := b.child(key)
		if i < 0 {
			return Loc{}, false, nil
		}
		next, err := r.blockAt(ctx, b.locs[i])
		if err == nil {
			err = checkChild(b, i, next)
		}
		if err != nil {
			return Loc{}, false, err
		}
		b = next
	}
	return Loc{}, false, corrupt("directory deeper than allowed")
}
