//ff:func feature=notation type=encoder control=selection
//ff:what parseField — one declared field from its token: a record of an allowed codec, or a raw field by ENUM name or raw words of its width
package notation

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// parseField encodes tok as field f. A record field's record must match one
// of the field's CODECs (any when none). A raw field is an ENUM name or raw
// words, exactly Width words unless it is the rest field.
func parseField(tok string, f codecspec.Field, ctx Context) ([]core.Word, error) {
	if f.Width == 0 {
		return recordField(tok, f, ctx)
	}
	ws, err := enumWords(tok, f.Enums)
	switch {
	case err != nil:
		return nil, err
	case f.Width != int(core.FieldRest) && len(ws) != f.Width:
		return nil, fmt.Errorf("%s is %d words, the field is %d", tok, len(ws), f.Width)
	default:
		return ws, nil
	}
}
