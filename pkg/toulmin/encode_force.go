//ff:func feature=toulmin type=encoder control=sequence
//ff:what encodeForce — append the F_FORCE attack force as a trimmed SIGNED-DECIMAL Value (ATTACKS only, [0, 1])
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// encodeForce appends p's attack force to out as a SIGNED-DECIMAL Value record
// with trailing zeros trimmed (0.5 is C=5, E=-1). It is called only when F_FORCE
// is set: the relation must be ATTACKS, Force must be present, and its value
// must lie in [0, 1]. An explicit 1 is kept (F_FORCE on, value 1).
func encodeForce(out []core.Word, p EdgePacket) ([]core.Word, error) {
	if p.Prop != PropAttacks {
		return nil, errors.New("geul/toulmin: F_FORCE is allowed only on ATTACKS")
	}
	if p.Force == nil {
		return nil, errors.New("geul/toulmin: F_FORCE set but Force is nil")
	}
	if *p.Force < 0 || *p.Force > vfixScale {
		return nil, errors.New("geul/toulmin: attack force out of range [0, 1]")
	}
	return PutValue(out, EncodeVFixValue(*p.Force))
}
