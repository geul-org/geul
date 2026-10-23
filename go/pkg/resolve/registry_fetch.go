//ff:func feature=resolve type=resolver control=sequence
//ff:what Registry.Fetch — always ErrNotFound until the registry API exists
package resolve

import "context"

// Fetch reports ErrNotFound: no Issuer is registered with a base address yet.
func (Registry) Fetch(context.Context, string) ([]byte, error) {
	return nil, ErrNotFound
}
