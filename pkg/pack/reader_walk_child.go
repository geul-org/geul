//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.walkChild — read internal child i, check it, and walk it one level deeper
package pack

import "context"

// walkChild descends into child i of b.
func (r *Reader) walkChild(ctx context.Context, b block, i int, lo, hi string, depth int, fn func(string, Loc) error) error {
	next, err := r.blockAt(ctx, b.locs[i])
	if err == nil {
		err = checkChild(b, i, next)
	}
	if err != nil {
		return err
	}
	return r.walk(ctx, next, lo, hi, depth+1, fn)
}
