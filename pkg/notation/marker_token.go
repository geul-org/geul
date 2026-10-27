//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what markerToken — the token of a named slot marker record (ABSENT, UNKNOWN, NONE); false for any other record
package notation

import "github.com/geul-org/gogeul/pkg/core"

// markerToken returns the markerTokens token of r when r is a named slot
// marker (core.IsSlotMarker, number 0~2).
func markerToken(r core.Record) (string, bool) {
	if !core.IsSlotMarker(r) {
		return "", false
	}
	for tok, w := range markerTokens {
		if core.Word(r.Codec) == w {
			return tok, true
		}
	}
	return "", false
}
