//ff:type feature=inspect type=model
//ff:what ImportDoc — Meta IMPORT-DOC: IMPORT ID and the DOC-POINTER Issuer, document ID and 2-word version ID
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// ImportDoc is Meta IMPORT-DOC (GEUL t002 OVERVIEW IMPORT): the IMPORT ID this
// document uses for the imported one, its Issuer and 4-word document ID, and
// the 2-word version ID it is pinned to (the publisher's sequence number in that
// document, GEUL t002 OVERVIEW 계보). Version 0 means the latest version.
type ImportDoc struct {
	ID      uint64
	Issuer  []core.Word
	Doc     []core.Word
	Version uint32
}
