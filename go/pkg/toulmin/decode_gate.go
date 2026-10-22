//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeGate — decode an operator word plus percent Value, enforcing reserved bits and range
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// DecodeGate decodes a Gate from the head of ws and returns the words consumed.
// The operator word's low 14 bits are reserved and must be 0; the percent must
// be an UNSIGNED-DECIMAL Value that is an exact integer in 0-100.
func DecodeGate(ws []core.Word) (Gate, int, error) {
	if len(ws) < 1 {
		return Gate{}, 0, errors.New("geul/toulmin: gate operator word missing")
	}
	if ws[0]&0x3FFF != 0 {
		return Gate{}, 0, errors.New("geul/toulmin: gate operator reserved bits must be 0")
	}
	r, n, err := core.DecodeRecord(ws[1:])
	if err != nil {
		return Gate{}, 0, err
	}
	if r.Codec != core.CodecUnsignedDecimal {
		return Gate{}, 0, errors.New("geul/toulmin: gate percent must be an UNSIGNED-DECIMAL Value")
	}
	e, c, err := core.DecodeDecimalValue(r)
	if err != nil {
		return Gate{}, 0, err
	}
	p, err := core.ScaleDecimal(e, c, 0, gateMaxPercent)
	if err != nil {
		return Gate{}, 0, err
	}
	return Gate{Op: GateOp(ws[0] >> 14), Percent: uint8(p)}, 1 + n, nil
}
