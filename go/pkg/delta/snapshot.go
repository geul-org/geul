//ff:type feature=delta type=model
//ff:what snapshot — a document split for diffing: HEAD records, and BODY Edges in order with their Edge Ids
package delta

import "github.com/geul-org/geul/go/pkg/core"

// snapshot is a document split for diffing (OVERVIEW 계보 diff): its HEAD
// records and its BODY Edges in order, each with its Edge Id.
type snapshot struct {
	head []core.Record
	body []core.Record
	ids  []uint64
}
