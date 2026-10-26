//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what deltaBaseOf — the base version a stored version file names in DELTA-BASE; false when it is a snapshot
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// deltaBaseOf reads the DELTA-BASE of document b's HEAD (meta/lineage/
// DELTA-BASE). It is false for a snapshot or a file that does not split.
func deltaBaseOf(b []byte) (uint32, bool) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return 0, false
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return 0, false
	}
	for _, r := range doc.Head {
		if v, err := core.DecodeDeltaBase(r); err == nil {
			return v, true
		}
	}
	return 0, false
}
