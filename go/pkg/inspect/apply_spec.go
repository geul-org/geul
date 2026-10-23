//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what applySpec — set each entry's codec spec, and an Edge's Edge Id when it has none
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// applySpec gives each entry the fields s declares for its codec. A declared
// Edge without an Edge Id from its typed content takes its leading LOCAL one.
func applySpec(entries []Entry, s codecspec.Spec) {
	for i := range entries {
		e := &entries[i]
		fs, ok := s.Of(e.Record)
		if !ok {
			continue
		}
		e.Spec = fs
		if e.Record.Kind != core.KindEdge || e.ID != 0 {
			continue
		}
		if r, _, err := core.DecodeRecord(e.Record.Payload); err == nil {
			e.ID, _ = core.DecodeLocalRef(r)
		}
	}
}
