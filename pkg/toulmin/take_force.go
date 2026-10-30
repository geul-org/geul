//ff:func feature=toulmin type=decoder control=sequence
//ff:what takeForce — read the F_FORCE attack force at *pos: ATTACKS only, a SIGNED-DECIMAL in [0, 1]
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// takeForce reads the attack force of a LINK at ws[*pos] and advances *pos. Only
// ATTACKS carries a force; on any other relation F_FORCE is an error. The value
// follows the TOULMIN numeric rule (SIGNED-DECIMAL, exact at four decimal places)
// and must lie in [0, 1].
func takeForce(ws []core.Word, pos *int, prop PropCode) (*VFix, error) {
	if prop != PropAttacks {
		return nil, errors.New("geul/toulmin: F_FORCE is allowed only on ATTACKS")
	}
	v, err := TakeVFix(ws, pos)
	if err != nil {
		return nil, err
	}
	if *v < 0 {
		return nil, errors.New("geul/toulmin: attack force out of range [0, 1]")
	}
	return v, nil
}
