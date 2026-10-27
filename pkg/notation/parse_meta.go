//ff:func feature=notation type=encoder control=selection
//ff:what ParseMeta — one line as a HEAD Meta record: its name and payload tokens, or the whole record as raw words
package notation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// ParseMeta reads `NAME token…` (a core Meta name, case ignored; the
// payload is the tokens read by parseGeneric) or `0x…`, a whole Meta record.
func ParseMeta(line string, ctx Context) (core.Record, error) {
	tokens, err := tokenize(line)
	switch {
	case err != nil:
		return core.Record{}, err
	case len(tokens) == 0:
		return core.Record{}, errors.New("geul/notation: no Meta name")
	case strings.HasPrefix(tokens[0], "0x") && len(tokens) == 1:
		r, err := parseRawRecord(tokens[0])
		if err == nil && r.Kind != core.KindMeta {
			err = fmt.Errorf("geul/notation: %s is not a Meta record", tokens[0])
		}
		return r, err
	}
	c, ok := codecOf(core.KindMeta, tokens[0])
	if !ok {
		return core.Record{}, fmt.Errorf("geul/notation: no Meta named %s", tokens[0])
	}
	payload, err := parseGeneric(tokens[1:], ctx)
	return core.Record{Kind: core.KindMeta, Codec: c, Payload: payload}, err
}
