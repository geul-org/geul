//ff:func feature=inspect type=decoder control=sequence
//ff:what parseDocName — Meta DOC-NAME as DocName; false when the Issuer width is not 1, 2, 4 or 16
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseDocName splits Meta DOC-NAME into the Issuer and the 4-word document
// ID. The Issuer width is the payload minus 4 and must be 1, 2, 4 or 16 words.
func parseDocName(r core.Record) (DocName, bool) {
	iw := len(r.Payload) - 4
	if iw != 1 && iw != 2 && iw != 4 && iw != 16 {
		return DocName{}, false
	}
	return DocName{Issuer: r.Payload[:iw], Doc: r.Payload[iw:]}, true
}
