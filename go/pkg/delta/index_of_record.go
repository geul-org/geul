//ff:func feature=delta type=helper control=iteration dimension=1
//ff:what indexOfRecord — the first index of a record equal to r, or -1
package delta

import "github.com/geul-org/geul/go/pkg/core"

// indexOfRecord finds r in rs by encoding (sameRecord).
func indexOfRecord(rs []core.Record, r core.Record) int {
	for i, x := range rs {
		if sameRecord(x, r) {
			return i
		}
	}
	return -1
}
