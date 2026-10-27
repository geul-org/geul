//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what quantityNumbers — why the first numeric record of a QUANTITY (amount or a present bound) is not canonical; "" when all are
package schema

import "github.com/geul-org/gogeul/pkg/core"

// quantityNumbers checks the amount and each bound that is not Absent with
// numberProblem and returns the first problem.
func quantityNumbers(q core.Quantity) string {
	for _, d := range []core.Record{q.Amount, q.Lower, q.Upper} {
		if core.IsAbsent(d) {
			continue
		}
		if p := numberProblem(d); p != "" {
			return p
		}
	}
	return ""
}
