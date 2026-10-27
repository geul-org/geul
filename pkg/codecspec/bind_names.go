//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what bindNames — the name each BIND puts on a LOCAL Subject, the last BIND winning
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// bindNames maps each LOCAL Subject of a BIND in body to its name. A later
// BIND replaces an earlier one (edge/BIND 대표 이름).
func bindNames(body []core.Record) map[uint64]string {
	names := map[uint64]string{}
	for _, r := range body {
		b, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if id, err := core.DecodeLocalRef(b.Subject); err == nil {
			names[id] = string(b.Name)
		}
	}
	return names
}
