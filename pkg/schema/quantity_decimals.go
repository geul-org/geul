//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what quantityDecimals — the first SIGNED-DECIMAL of a QUANTITY (amount or a present bound) whose coefficient is not minimal
package schema

import "github.com/geul-org/gogeul/pkg/core"

// quantityDecimals decodes the amount and each bound that is not Absent and
// returns the first error (a non-minimal coefficient).
func quantityDecimals(q core.Quantity) error {
	for _, d := range []core.Record{q.Amount, q.Lower, q.Upper} {
		if core.IsAbsent(d) {
			continue
		}
		if _, _, err := core.DecodeSignedDecimalBig(d); err != nil {
			return err
		}
	}
	return nil
}
