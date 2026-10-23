//ff:func feature=inspect type=decoder control=sequence
//ff:what parseDocID — Meta DOC-ID as DocID; false when it is not one EXTERNAL-ISSUER-*(Issuer, document, 0)
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseDocID reads Meta DOC-ID (core.DecodeDocID) into the Issuer and the
// 4-word document ID.
func parseDocID(r core.Record) (DocID, bool) {
	issuer, doc, err := core.DecodeDocID(r)
	if err != nil {
		return DocID{}, false
	}
	return DocID{Issuer: issuer, Doc: doc}, true
}
