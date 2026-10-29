//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.hashPath — the `h/` path a record serves, from its bytes' content hash (§4.5, §8.1)
package pack

import (
	"context"
)

// hashPath reads the record at l and names it by its content hash.
func (r *Reader) hashPath(ctx context.Context, l Loc) (string, error) {
	b, err := r.record(ctx, l)
	if err != nil {
		return "", err
	}
	h, ok := contentHash(b)
	if !ok {
		return "", corrupt("h/ record is not a GEUL document")
	}
	return hPath(string(h[:])), nil
}
