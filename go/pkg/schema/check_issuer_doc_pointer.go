//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerDocPointer — R5: ISSUER-POINTER·DOC-POINTER Quantifier Identified, Issuer width 1/2/4/16, DOC-POINTER document ID not 0
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkIssuerDocPointer applies R5 to an ISSUER-POINTER or DOC-POINTER
// Reference whose 4-bit codec is c. The Issuer width is what the Length leaves
// after the fixed fields: nothing for ISSUER-POINTER, the 4-word document ID
// and 2-word version ID for DOC-POINTER.
func checkIssuerDocPointer(r rec, c uint8) []Fact {
	var fs []Fact
	if q := r.R.Codec & 7; q != 0 {
		fs = append(fs, factAt("R5", Error, r, "Quantifier", "000 (Identified)", fmt.Sprintf("%03b", q)))
	}
	p := r.R.Payload
	fixed := 0
	if c == core.RefDocPointer {
		fixed = 6
	}
	w := len(p) - fixed
	if w != 1 && w != 2 && w != 4 && w != 16 {
		return append(fs, factAt("R5", Error, r, "Issuer", "1/2/4/16 WORD", fmtWords(w)))
	}
	if c == core.RefDocPointer && p[w]|p[w+1]|p[w+2]|p[w+3] == 0 {
		fs = append(fs, factAt("R5", Error, r, "문서 ID", "0이 아님 (Issuer 자신은 ISSUER-POINTER)", "0"))
	}
	return fs
}
