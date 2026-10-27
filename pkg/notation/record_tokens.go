//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what recordTokens — words as one token per record when they are wholly records that each have a short form
package notation

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// recordTokens decodes ws record by record. It is false when a record does
// not decode or would be written as raw words (formatRecord), so a part is
// never split into raw pieces.
func recordTokens(ws []core.Word) ([]string, bool) {
	var out []string
	for len(ws) > 0 {
		r, n, err := core.DecodeRecord(ws)
		if err != nil {
			return nil, false
		}
		s := formatRecord(r)
		if strings.HasPrefix(s, "0x") && !strings.Contains(s, "/") {
			return nil, false
		}
		out = append(out, s)
		ws = ws[n:]
	}
	return out, len(out) > 0
}
