//ff:func feature=schema type=rule control=sequence
//ff:what minimalSigned — a two's-complement integer in the fewest words
package schema

import "github.com/geul-org/gogeul/pkg/core"

// minimalSigned reports whether ws is a non-empty two's-complement integer whose
// leading 0x0000/0xFFFF word, if any, is needed to carry the sign.
func minimalSigned(ws []core.Word) bool {
	if len(ws) < 2 {
		return len(ws) == 1
	}
	pos := ws[0] == 0 && ws[1]&0x8000 == 0
	neg := ws[0] == 0xFFFF && ws[1]&0x8000 != 0
	return !pos && !neg
}
