//ff:type feature=inspect type=model
//ff:what ImportDoc — Meta IMPORT-DOC: IMPORT ID, Issuer, document ID and optional pinned version ID
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// ImportDoc is Meta IMPORT-DOC (GEUL t002 OVERVIEW IMPORT): the IMPORT ID this
// document uses for the imported one, its Issuer and 4-word document ID, and
// the 16-word version ID when the import is pinned (nil otherwise).
type ImportDoc struct {
	ID      uint64
	Issuer  []core.Word
	Doc     []core.Word
	Version []core.Word
}
