//ff:func feature=notation type=encoder control=selection
//ff:what ParseEdge — one line as an Edge: optional #label, the Edge's name, its fields (by codec spec when declared)
package notation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseEdge reads `[#label] NAME token…` (or `[#label] 0x…`, a whole Edge as
// raw words, which keeps its own Edge Id as the label). NAME is a core Edge
// name (case ignored). The fields after the Edge Id follow ctx.Spec when it
// declares the codec (parseFields), else are read token by token
// (parseGeneric).
func ParseEdge(line string, ctx Context) (Edge, error) {
	tokens, err := tokenize(line)
	if err != nil {
		return Edge{}, err
	}
	var e Edge
	if len(tokens) > 0 && strings.HasPrefix(tokens[0], "#") {
		e.Label, err = strconv.ParseUint(tokens[0][1:], 10, 64)
		if err != nil {
			return Edge{}, fmt.Errorf("geul/notation: label %q is not #n", tokens[0])
		}
		tokens = tokens[1:]
	}
	switch {
	case len(tokens) == 0:
		return Edge{}, errors.New("geul/notation: no Edge name")
	case strings.HasPrefix(tokens[0], "0x") && len(tokens) == 1:
		return rawEdge(tokens[0])
	}
	return namedEdge(e, tokens, ctx)
}
