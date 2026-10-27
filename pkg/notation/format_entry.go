//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what FormatEntry — one interpreted record as a line: #id NAME tokens for an Edge, NAME tokens otherwise, raw words when it has no readable name
package notation

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// FormatEntry writes e by its layout (inspect.Layout): an Edge as its Edge
// Id `#n`, its codec name and one token per field after the Edge Id; any
// other record as its codec name and its fields. A record whose name does not
// read back to its codec (codecOf), that carries an Issuer or format number,
// or an Edge that does not open with a LOCAL Edge Id is its raw words.
func FormatEntry(e inspect.Entry) string {
	raw := hexToken(e.Words)
	c, ok := codecOf(e.Record.Kind, e.Name)
	if !ok || c != e.Record.Codec || len(e.Record.Issuer) > 0 || e.Record.Format != 0 {
		return raw
	}
	parts := payloadOnly(inspect.Layout(e))
	tokens := []string{e.Name}
	if e.Record.Kind == core.KindEdge {
		id, ok := edgeIDToken(parts)
		if !ok {
			return raw
		}
		tokens, parts = []string{id, e.Name}, parts[1:]
	}
	for _, p := range parts {
		tokens = append(tokens, partTokens(p)...)
	}
	return strings.Join(tokens, " ")
}
