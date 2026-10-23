//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what checkVersionList — P1~P4 over a version list document
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// checkVersionList applies the version list rules: no IMPORT, only version list
// records (P1), Edge Ids 1, 2, … (P2), and the per-document rules (P3, P4).
func checkVersionList(parts docParts) []Fact {
	imports, _ := importsOf(parts.Head)
	var fs []Fact
	if len(imports) > 0 {
		fs = append(fs, Fact{Rule: "P1", Severity: Error, Expected: "IMPORT-DOC-* 없음", Actual: fmt.Sprintf("%d개", len(imports))})
	}
	st := map[string]*docVersions{}
	for i, r := range parts.Body {
		p, err := core.DecodePublication(r.R)
		if err != nil {
			fs = append(fs, factAt("P1", Error, r, "", "PUBLICATION·DEPRECATION·TOMBSTONE", fmt.Sprintf("kind %d codec 0x%02X", r.R.Kind, r.R.Codec)))
			continue
		}
		if p.EdgeID != uint64(i+1) {
			fs = append(fs, factAt("P2", Error, r, "Edge Id", fmt.Sprintf("%d", i+1), fmt.Sprintf("%d", p.EdgeID)))
		}
		fs = append(fs, applyPublication(st, r, p)...)
	}
	return fs
}
