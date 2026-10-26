//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what headChanges — HEAD records of base missing from next (removed) and of next missing from base (added), each matched once
package delta

import "github.com/geul-org/geul/go/pkg/core"

// headChanges matches equal records of base and next one to one and reports
// the rest: base's as '-', in order, then next's as '+'.
func headChanges(base, next []core.Record) []Change {
	left := append([]core.Record{}, next...)
	var out []Change
	for _, r := range base {
		if i := indexOfRecord(left, r); i >= 0 {
			left = append(left[:i], left[i+1:]...)
			continue
		}
		out = append(out, Change{Op: '-', Head: true, Old: r})
	}
	for _, r := range left {
		out = append(out, Change{Op: '+', Head: true, New: r})
	}
	return out
}
