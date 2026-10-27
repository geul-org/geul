//ff:func feature=schema type=rule control=sequence
//ff:what minimalUnsigned — an unsigned integer in the fewest words (no leading 0x0000 unless the value is 0)
package schema

import "github.com/geul-org/gogeul/pkg/core"

// minimalUnsigned reports whether ws is a non-empty unsigned integer without a
// removable leading 0x0000 word.
func minimalUnsigned(ws []core.Word) bool {
	return len(ws) == 1 || (len(ws) > 1 && ws[0] != 0)
}
