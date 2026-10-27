//ff:func feature=schema type=rule control=selection
//ff:what finiteFloat — a BINARY-FLOAT Payload of an IEEE width (1/2/4/8 WORD) whose value is finite (not NaN or ±Infinity)
package schema

import "github.com/geul-org/gogeul/pkg/core"

// finiteFloat reports whether ws is a binary16/32/64/128 interchange value
// (1/2/4/8 WORDs) that is finite: its exponent field, in the first WORD after
// the sign bit, is not all ones (NaN and ±Infinity are).
func finiteFloat(ws []core.Word) bool {
	var mask core.Word
	switch len(ws) {
	case 1:
		mask = 0x7C00
	case 2:
		mask = 0x7F80
	case 4:
		mask = 0x7FF0
	case 8:
		mask = 0x7FFF
	default:
		return false
	}
	return ws[0]&mask != mask
}
