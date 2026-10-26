//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what parseGeneric — payload words from tokens with no declared fields: raw words as they are, every other token a record
package notation

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseGeneric concatenates the tokens: `0x…` (without `/`) is raw words,
// anything else one record (parseRecord).
func parseGeneric(tokens []string, ctx Context) ([]core.Word, error) {
	var out []core.Word
	for _, t := range tokens {
		var ws []core.Word
		var err error
		if strings.HasPrefix(t, "0x") && !strings.Contains(t, "/") {
			ws, err = parseHex(t)
		} else {
			ws, err = recordField(t, recordFieldAny, ctx)
		}
		if err != nil {
			return nil, err
		}
		out = append(out, ws...)
	}
	return out, nil
}
