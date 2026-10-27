//ff:type feature=resolve type=model
//ff:what Key — a document address: Issuer, 4-word document ID and 2-word version ID (0 = latest)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// Key addresses one document (GEUL-ORG 공표 주소): the Issuer (1, 2, 4 or 16
// words), the 4-word document ID and the version ID, the publisher's sequence
// number in that document. Version 0 asks for the latest version.
type Key struct {
	Issuer  []core.Word
	Doc     []core.Word
	Version uint32
}
