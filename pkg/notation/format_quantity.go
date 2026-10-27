//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what formatQuantity — QUANTITY(amount, unit, lower, upper) with each part as a token and ABSENT for a part left out
package notation

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatQuantity writes the four parts of a QUANTITY Value (formatRecord).
func formatQuantity(r core.Record) (string, bool) {
	q, err := core.DecodeQuantity(r)
	if err != nil {
		return "", false
	}
	parts := []string{}
	for _, p := range []core.Record{q.Amount, q.Unit, q.Lower, q.Upper} {
		parts = append(parts, formatRecord(p))
	}
	return "QUANTITY(" + strings.Join(parts, ", ") + ")", true
}
