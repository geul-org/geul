//ff:func feature=inspect type=decoder control=selection
//ff:what binaryFloat — a BINARY-FLOAT Payload of 1, 2 or 4 WORDs (binary16/32/64) as float64; binary128 and other widths fail
package inspect

import (
	"errors"
	"math"

	"github.com/geul-org/geul/go/pkg/core"
)

// binaryFloat reads a BINARY-FLOAT Payload (value/numeric/BINARY-FLOAT):
// binary16, binary32 or binary64 by 1, 2 or 4 WORDs, exactly as float64.
// binary128 (8 WORDs) does not fit a float64 and fails, as does any other
// width.
func binaryFloat(ws []core.Word) (float64, error) {
	switch len(ws) {
	case 1:
		return half(uint16(ws[0])), nil
	case 2:
		return float64(math.Float32frombits(uint32(ws[0])<<16 | uint32(ws[1]))), nil
	case 4:
		return math.Float64frombits(uint64(ws[0])<<48 | uint64(ws[1])<<32 | uint64(ws[2])<<16 | uint64(ws[3])), nil
	default:
		return 0, errors.New("geul/inspect: BINARY-FLOAT wider than binary64 or of no IEEE width")
	}
}
