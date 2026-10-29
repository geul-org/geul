//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.walkEntry — a leaf entry goes to fn, an internal entry is walked one level down
package pack

import "context"

// walkEntry handles entry i of b during a walk.
func (r *Reader) walkEntry(ctx context.Context, b block, i int, lo, hi string, depth int, fn func(string, Loc) error) error {
	if isLeaf(b.kind) {
		return fn(b.keys[i], b.locs[i])
	}
	return r.walkChild(ctx, b, i, lo, hi, depth, fn)
}
