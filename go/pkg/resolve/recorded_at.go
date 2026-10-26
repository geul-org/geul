//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what recordedAt — the RECORDED-AT time of a version list HEAD; zero when absent or unreadable
package resolve

import (
	"time"

	"github.com/geul-org/geul/go/pkg/core"
)

// recordedAt reads the RECORDED-AT Meta (0x25) of head.
func recordedAt(head []core.Record) time.Time {
	for _, r := range head {
		if r.Kind != core.KindMeta || r.Codec != 0x25 {
			continue
		}
		tv, _ := metaTime(r)
		return tv.Time
	}
	return time.Time{}
}
