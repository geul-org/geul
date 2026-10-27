//ff:type feature=delta type=model
//ff:what Change — one difference between two versions: a HEAD record or an Edge (by Edge Id) added, removed or changed
package delta

import "github.com/geul-org/gogeul/pkg/core"

// Change is one difference found by Changes. Op is '+' (added), '-'
// (removed) or '~' (same Edge Id, other content). Head marks a HEAD record,
// which has no ID and is only added or removed. Old and New are the record
// before and after (the zero Record where there is none).
type Change struct {
	Op   byte
	Head bool
	ID   uint64
	Old  core.Record
	New  core.Record
}
