//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerRef — R4: EXTERNAL-ISSUER Issuer width, document ID and target
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkIssuerRef applies R4 (and R2 to the target) to an EXTERNAL-ISSUER-*
// Reference whose 4-bit codec is c.
func checkIssuerRef(r rec, c uint8) []Fact {
	w := core.ImportIssuerWidth(core.MetaImportDoc1Word + c&3)
	p := r.R.Payload
	if len(p) < w+5 {
		return []Fact{factAt("R4", Error, r, "Payload", fmt.Sprintf("Issuer %d WORD + 문서 ID 4 WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	return checkTarget(r, p[w+4:])
}
