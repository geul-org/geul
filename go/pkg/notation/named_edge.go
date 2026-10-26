//ff:func feature=notation type=encoder control=selection
//ff:what namedEdge — the codec and payload of NAME token… for an Edge read by ParseEdge
package notation

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// namedEdge fills e from tokens[0], an Edge name, and its field tokens.
func namedEdge(e Edge, tokens []string, ctx Context) (Edge, error) {
	c, ok := codecOf(core.KindEdge, tokens[0])
	if !ok {
		return Edge{}, fmt.Errorf("geul/notation: no Edge named %s", tokens[0])
	}
	e.Codec = c
	var err error
	if fields, ok := ctx.Spec.Of(core.Record{Kind: core.KindEdge, Codec: c}); ok {
		e.Rest, err = parseFields(tokens[1:], fields, ctx)
	} else {
		e.Rest, err = parseGeneric(tokens[1:], ctx)
	}
	switch {
	case err != nil:
		return Edge{}, err
	default:
		return e, nil
	}
}
