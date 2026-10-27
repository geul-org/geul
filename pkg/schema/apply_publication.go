//ff:func feature=schema type=rule control=selection
//ff:what applyPublication — P3, P4: one version list record against the state of its document
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// applyPublication checks one decoded version list record against st (P3, P4)
// and records it.
func applyPublication(st map[string]*docVersions, r rec, p core.Publication) []Fact {
	key := core.FormatID(p.Doc)
	d := st[key]
	if d == nil {
		d = &docVersions{hashes: map[[32]byte]bool{}}
		st[key] = d
	}
	switch {
	case p.Codec != core.EdgePublication && p.Version > d.last:
		return []Fact{factAt("P4", Error, r, "버전 ID", "앞에 공표된 버전", fmt.Sprintf("%s v%d (미공표)", key, p.Version))}
	case p.Codec != core.EdgePublication:
		return nil
	case p.Version != d.last+1:
		want := d.last + 1
		d.last = p.Version
		return []Fact{factAt("P3", Error, r, "버전 ID", fmt.Sprintf("v%d", want), fmt.Sprintf("v%d", p.Version))}
	case d.hashes[p.Hash]:
		d.last = p.Version
		return []Fact{factAt("P3", Error, r, "내용 해시", "처음 공표하는 내용", "같은 내용 재공표")}
	default:
		d.last = p.Version
		d.hashes[p.Hash] = true
		return nil
	}
}
