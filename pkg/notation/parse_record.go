//ff:func feature=notation type=encoder control=selection
//ff:what parseRecord — one token as one embedded record: a slot marker (ABSENT, UNKNOWN, NONE), raw words, #n/Qn/Pn, quoted text, <URI>, NAME(…), a document address or a decimal
package notation

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseRecord reads a record token (package doc). `ABSENT`, `UNKNOWN` and `NONE`
// are the slot markers ABSENT, UNKNOWN-VALUE and NO-VALUE (markerTokens);
// whether a field takes one is recordField's check.
func parseRecord(tok string, ctx Context) (core.Record, error) {
	switch {
	case tok == "":
		return core.Record{}, fmt.Errorf("geul/notation: empty token")
	case markerTokens[tok] != 0 || tok == "ABSENT":
		r, _, _ := core.DecodeRecord([]core.Word{markerTokens[tok]})
		return r, nil
	case strings.HasPrefix(tok, "0x") && !strings.Contains(tok, "/"):
		return parseRawRecord(tok)
	case tok[0] == '#' || (tok[0] == 'Q' || tok[0] == 'P') && len(tok) > 1 && tok[1] >= '0' && tok[1] <= '9':
		return parseNumbered(tok, ctx)
	case tok[0] == '"':
		return parseText(tok, ctx)
	case tok[0] == '<' && strings.HasSuffix(tok, ">"):
		return core.EncodeBytesValue(codecURIEven, []byte(tok[1:len(tok)-1])), nil
	case strings.HasSuffix(tok, ")") && strings.Contains(tok, "("):
		return parseCall(tok, ctx)
	case strings.Contains(tok, "/"):
		return parseAddress(tok, ctx)
	case strings.ContainsAny(tok[:1], "+-.0123456789"):
		return parseDecimalValue(core.CodecSignedDecimal, tok)
	default:
		return core.Record{}, fmt.Errorf("geul/notation: cannot read %q", tok)
	}
}
