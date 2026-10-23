//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerRef — R4: EXTERNAL-ISSUER Issuer width, document ID and target; document 0 only with target 0
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkIssuerRef applies R4 (and R2 to the target) to an EXTERNAL-ISSUER-*
// Reference whose 4-bit codec is c. Document ID 0 (the Issuer's version list
// document) admits only target ID 0, the document itself.
func checkIssuerRef(r rec, c uint8) []Fact {
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := r.R.Payload
	if len(p) < w+5 {
		return []Fact{factAt("R4", Error, r, "Payload", fmt.Sprintf("Issuer %d WORD + 문서 ID 4 WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	fs := checkTarget(r, p[w+4:])
	if p[w]|p[w+1]|p[w+2]|p[w+3] == 0 && (len(p) != w+5 || p[w+4] != 0) {
		fs = append(fs, factAt("R4", Error, r, "대상 ID", "문서 ID 0이면 대상 ID 0 (Issuer 대표)", fmtWords(len(p)-w-4)))
	}
	return fs
}
