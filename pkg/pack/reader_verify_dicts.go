//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what Reader.verifyDicts — every DICTS dictionary loads: bytes match their SHA-256 and make a decoder under their ID
package pack

import "context"

// verifyDicts loads each dictionary once.
func (r *Reader) verifyDicts(ctx context.Context) error {
	for _, d := range r.m.dicts {
		if _, err := r.decoder(ctx, d.id); err != nil {
			return err
		}
	}
	return nil
}
