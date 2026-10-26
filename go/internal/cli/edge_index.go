//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what edgeIndex — the position in a BODY of the Edge with an Edge Id; -1 when none
package cli

import "github.com/geul-org/geul/go/pkg/core"

// edgeIndex finds the Edge whose leading LOCAL is id.
func edgeIndex(body []core.Record, id uint64) int {
	for i, r := range body {
		if r.Kind != core.KindEdge {
			continue
		}
		lead, _, err := core.DecodeRecord(r.Payload)
		if n, lerr := core.DecodeLocalRef(lead); err == nil && lerr == nil && n == id {
			return i
		}
	}
	return -1
}
