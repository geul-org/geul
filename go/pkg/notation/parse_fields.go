//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what parseFields — the payload words of declared fields from tokens, one token per field (a rest field may be left out)
package notation

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// parseFields reads tokens against fields in order (parseField). A rest field
// (core.FieldRest) takes a token only when there are more tokens than the
// fields after it. Tokens left over or missing are an error.
func parseFields(tokens []string, fields []codecspec.Field, ctx Context) ([]core.Word, error) {
	var out []core.Word
	i := 0
	for j, f := range fields {
		if f.Width == int(core.FieldRest) && len(tokens)-i <= len(fields)-j-1 {
			continue
		}
		if i >= len(tokens) {
			return nil, fmt.Errorf("geul/notation: field %s is missing", f.Name)
		}
		ws, err := parseField(tokens[i], f, ctx)
		if err != nil {
			return nil, fmt.Errorf("geul/notation: field %s: %w", f.Name, err)
		}
		out = append(out, ws...)
		i++
	}
	if i != len(tokens) {
		return nil, fmt.Errorf("geul/notation: %d tokens for %d fields", len(tokens), len(fields))
	}
	return out, nil
}
