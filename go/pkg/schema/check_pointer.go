//ff:func feature=schema type=rule control=sequence
//ff:what checkPointer — R3: EXTERNAL-POINTER IMPORT ID width, narrowest codec and its IMPORT-DOC
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkPointer applies R3 (and R2 to the target) to an EXTERNAL-POINTER-*
// Reference whose 4-bit codec is c. imports holds the IMPORT IDs of the HEAD.
func checkPointer(r rec, c uint8, imports map[uint64]bool) []Fact {
	w := core.WidthWords(c)
	p := r.R.Payload
	if len(p) < w+1 {
		return []Fact{factAt("R3", Error, r, "Payload", fmt.Sprintf("IMPORT ID %d WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	var fs []Fact
	id, err := core.DecodeID(p[:w], w)
	if err == nil && w > core.MinWidth(id) {
		fs = append(fs, factAt("R3", Error, r, "IMPORT ID", fmt.Sprintf("EXTERNAL-POINTER-%dWORD", core.MinWidth(id)), fmt.Sprintf("EXTERNAL-POINTER-%dWORD", w)))
	}
	if err != nil || !imports[id] {
		fs = append(fs, factAt("R3", Error, r, "IMPORT ID", "HEAD의 IMPORT-DOC-*이 배정한 IMPORT ID", fmt.Sprintf("%d", id)))
	}
	return append(fs, checkTarget(r, p[w:])...)
}
