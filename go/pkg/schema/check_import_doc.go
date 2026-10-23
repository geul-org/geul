//ff:func feature=schema type=rule control=sequence
//ff:what checkImportDoc — M3: IMPORT-DOC layout and a minimal IMPORT ID
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkImportDoc applies M3 to one IMPORT-DOC-*: IMPORT ID (minimal, 1 WORD or
// more), Issuer of the codec's width, 4-word document ID, 2-word version ID.
func checkImportDoc(r rec) []Fact {
	w := core.ImportIssuerWidth(r.R.Codec)
	n := len(r.R.Payload) - w - 6
	if n < 1 {
		return []Fact{factAt("M3", Error, r, "Payload", fmt.Sprintf("IMPORT ID + Issuer %d WORD + 문서 ID 4 + 버전 ID 2", w), fmtWords(len(r.R.Payload)))}
	}
	if !minimalUnsigned(r.R.Payload[:n]) {
		return []Fact{factAt("M3", Error, r, "IMPORT ID", "최소 WORD", fmtWords(n))}
	}
	return nil
}
