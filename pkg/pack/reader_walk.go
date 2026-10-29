//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.walk — visit a subtree's leaf entries in [lo, hi) in key order, descending into overlapping children
package pack

import "context"

// walk calls fn for every leaf entry under b with lo ≤ key < hi (hi "" is
// unbounded), deletion markers included.
func (r *Reader) walk(ctx context.Context, b block, lo, hi string, depth int, fn func(string, Loc) error) error {
	if depth >= maxDepth {
		return corrupt("directory deeper than allowed")
	}
	for i := range b.keys {
		if !overlaps(b, i, lo, hi) {
			continue
		}
		if err := r.walkEntry(ctx, b, i, lo, hi, depth, fn); err != nil {
			return err
		}
	}
	return nil
}
