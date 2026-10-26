//ff:func feature=notation type=decoder control=sequence
//ff:what formatNumbered — LOCAL as #n, WIKIDATA-QID as Qn, WIKIDATA-PID as Pn, from a minimal unsigned target ID
package notation

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// formatNumbered writes the target ID of Reference codec c (CCCC). An ID
// that is not minimal or is past 64 bits is false, so it stays raw.
func formatNumbered(c uint8, payload []core.Word) (string, bool) {
	n, err := core.DecodeUintValue(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: payload})
	if err != nil {
		return "", false
	}
	prefix := map[uint8]string{core.RefLocal: "#", core.RefWikidataQID: "Q", core.RefWikidataPID: "P"}[c]
	return prefix + strconv.FormatUint(n, 10), true
}
