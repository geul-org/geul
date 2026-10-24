//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerRef — R4: EXTERNAL-ISSUER Issuer width, document ID not 0 and a target ID not 0
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkIssuerRef applies R4 (and R2 to the target) to an EXTERNAL-ISSUER-*
// Reference whose 4-bit codec is c. It points at a LOCAL ID of an Issuer's
// document: the Issuer itself is ISSUER-POINTER and the document itself
// DOC-POINTER, so document ID 0 and target ID 0 are invalid.
func checkIssuerRef(r rec, c uint8) []Fact {
	w := core.ExternalIssuerWidth(c)
	p := r.R.Payload
	if len(p) < w+5 {
		return []Fact{factAt("R4", Error, r, "Payload", fmt.Sprintf("Issuer %d WORD + 문서 ID 4 WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	fs := checkTarget(r, p[w+4:])
	if p[w]|p[w+1]|p[w+2]|p[w+3] == 0 {
		fs = append(fs, factAt("R4", Error, r, "문서 ID", "0이 아님 (Issuer 자신은 ISSUER-POINTER)", "0"))
	}
	if len(p) == w+5 && p[w+4] == 0 {
		fs = append(fs, factAt("R4", Error, r, "대상 ID", "0이 아님 (문서 자신은 DOC-POINTER)", "0"))
	}
	return fs
}
