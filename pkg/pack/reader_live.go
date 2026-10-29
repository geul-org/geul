//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.live — the live keys in [lo, hi) across layers: newest layer wins, deletion markers dropped
package pack

import "context"

// live merges the layers' entries in [lo, hi) and removes deleted keys.
func (r *Reader) live(ctx context.Context, ls []layer, lo, hi string) (keySet, error) {
	s := keySet{}
	for _, l := range ls {
		if err := r.walk(ctx, l.blk, lo, hi, 0, s.keep); err != nil {
			return nil, err
		}
	}
	for k, l := range s {
		if l.Enc == encGone {
			delete(s, k)
		}
	}
	return s, nil
}
