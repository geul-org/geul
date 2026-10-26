//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what timeOf — when BODY record j of a version list was first recorded: the RECORDED-AT of the smallest list holding it
package resolve

import "time"

// timeOf returns the time of the first list in ts (any order) that holds more
// than j records, so record j (0-based) was added by it.
func timeOf(ts []listTime, j int) time.Time {
	best := -1
	var at time.Time
	for _, t := range ts {
		if t.n > j && (best < 0 || t.n < best) {
			best, at = t.n, t.at
		}
	}
	return at
}
