//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeGate — encode a Gate as an operator word plus an UNSIGNED-DECIMAL percent Value
package toulmin

import (
	"errors"

	"github.com/geul-org/geul/go/pkg/core"
)

// EncodeGate encodes g as the operator word (op in bits 15-14, reserved bits 0)
// followed by the percent as an UNSIGNED-DECIMAL Value record (E=0). A percent
// above 100 is an error rather than a silent clamp.
func EncodeGate(g Gate) ([]core.Word, error) {
	if g.Percent > gateMaxPercent {
		return nil, errors.New("geul/toulmin: gate percent exceeds 100")
	}
	r, _ := core.EncodeDecimalValue(false, 0, int64(g.Percent))
	pw, _ := core.EncodeRecord(r)
	return append([]core.Word{core.Word(uint16(g.Op) << 14)}, pw...), nil
}
