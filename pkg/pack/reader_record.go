//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.record — a record's original bytes: one range read, then unframe by encoding
package pack

import "context"

// record reads and restores the record at l.
func (r *Reader) record(ctx context.Context, l Loc) ([]byte, error) {
	b, err := r.raw(ctx, kindData, l)
	if err != nil {
		return nil, err
	}
	return r.unstore(ctx, l.Enc, b, false)
}
