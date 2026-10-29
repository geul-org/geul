//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what Reader.VerifyObjects — every object in OBJECTS: SHA-256 = name, header = OBJECTS entry (§8.1)
package pack

import "context"

// VerifyObjects checks all objects; do it once after receiving a pack set.
func (r *Reader) VerifyObjects(ctx context.Context) error {
	for i := range r.m.objects {
		if err := r.verifyObject(ctx, i); err != nil {
			return err
		}
	}
	return nil
}
