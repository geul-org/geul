//ff:func feature=pack type=rule control=sequence
//ff:what Reader.Verify — objects hash-checked, inventories replayed (h/ content hashes recomputed) and equal to the index
package pack

import "context"

// Verify runs the full integrity check of a generation: VerifyObjects, then
// State (which recomputes every shared or h/-only record's content hash)
// compared with the layers and NAMED.
func (r *Reader) Verify(ctx context.Context) error {
	if err := r.VerifyObjects(ctx); err != nil {
		return err
	}
	st, err := r.State(ctx)
	if err != nil {
		return err
	}
	idx, err := r.entries(ctx)
	if err != nil {
		return err
	}
	return sameState(st, idx)
}
