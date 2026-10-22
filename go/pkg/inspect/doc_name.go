//ff:type feature=inspect type=model
//ff:what DocName — Meta DOC-NAME: the Issuer and the 4-word document ID
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// DocName is Meta DOC-NAME: the Issuer (1, 2, 4 or 16 words) and the 4-word
// document ID it issued.
type DocName struct {
	Issuer []core.Word
	Doc    []core.Word
}
