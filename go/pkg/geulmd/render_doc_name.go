//ff:func feature=geulmd type=encoder control=sequence
//ff:what renderDocName — `DOC-NAME <issuer> <doc id>` with digit counts that keep the widths
package geulmd

import "github.com/geul-org/geul/go/pkg/core"

// renderDocName renders Meta DOC-NAME as the Issuer and the 4-word document ID.
// The Issuer width is the payload minus 4 and must be 1, 2, 4 or 16 words.
func renderDocName(r core.Record) (string, bool) {
	iw := len(r.Payload) - 4
	if iw != 1 && iw != 2 && iw != 4 && iw != 16 {
		return "", false
	}
	return "DOC-NAME " + hexID(r.Payload[:iw]) + " " + hexID(r.Payload[iw:]), true
}
