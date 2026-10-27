//ff:func feature=notation type=encoder control=sequence
//ff:what parseCall — NAME(args) as the Value or Reference of that codec name
package notation

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseCall reads `NAME(arg, …)`: a byte-string Value from one quoted
// string, the decimals and integers from one number, the time codecs, QUANTITY
// from four record tokens, the coordinates (parseCoordinate), the digests from raw words, ISSUER-POINTER from an
// Issuer.
func parseCall(tok string, ctx Context) (core.Record, error) {
	i := strings.Index(tok, "(")
	name, args := strings.ToUpper(tok[:i]), splitArgs(tok[i+1:len(tok)-1])
	if name == "QUANTITY" {
		return parseQuantity(args, ctx)
	}
	if name == "COORDINATE-EARTH" || name == "COORDINATE-GEO" {
		return parseCoordinate(args, name == "COORDINATE-GEO", ctx)
	}
	if len(args) != 1 {
		return core.Record{}, fmt.Errorf("geul/notation: %s takes one argument", name)
	}
	if name == "ISSUER-POINTER" {
		k, err := lookupOf(ctx)(args[0] + "/0x0000000000000000")
		if err != nil {
			return core.Record{}, err
		}
		return core.EncodeIssuerPointer(k.Issuer)
	}
	c, ok := codecOf(core.KindValue, name)
	if !ok {
		return core.Record{}, fmt.Errorf("geul/notation: no Value or Reference named %s", name)
	}
	return parseValueArg(c, args[0])
}
