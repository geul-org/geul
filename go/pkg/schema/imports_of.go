//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what importsOf — the IMPORT IDs of the HEAD, reporting duplicates (M3)
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// importsOf collects the IMPORT IDs assigned by the HEAD's IMPORT-DOC-* records.
// An IMPORT ID assigned twice is an M3 violation.
func importsOf(head []rec) (map[uint64]bool, []Fact) {
	ids := map[uint64]bool{}
	var fs []Fact
	for _, r := range head {
		id, ok := importID(r)
		if core.ImportIssuerWidth(r.R.Codec) == 0 || !ok {
			continue
		}
		if ids[id] {
			fs = append(fs, factAt("M3", Error, r, "IMPORT ID", "문서 안에서 한 번만 배정", fmt.Sprintf("%d 중복", id)))
		}
		ids[id] = true
	}
	return ids, fs
}
