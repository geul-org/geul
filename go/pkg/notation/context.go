//ff:type feature=notation type=model
//ff:what Context — what reading a line needs: codec specs for field layouts, a Lookup for names, a map for #labels
package notation

import "github.com/geul-org/geul/go/pkg/codecspec"

// Context configures reading. Spec gives the fields of declared codecs (a
// codec without one is read token by token). Lookup resolves document names.
// Local, when set, maps every `#n` written in a token to the LOCAL ID to
// encode (labels of a batch, geul add -i).
type Context struct {
	Spec   codecspec.Spec
	Lookup Lookup
	Local  func(n uint64) uint64
}
