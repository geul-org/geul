//ff:func feature=notation type=encoder control=selection
//ff:what recordField — a record field's words: the token's record, refused when the field's CODECs do not allow its codec
package notation

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// recordField reads tok (parseRecord) and checks it against f.Codecs.
func recordField(tok string, f codecspec.Field, ctx Context) ([]core.Word, error) {
	r, err := parseRecord(tok, ctx)
	switch {
	case err != nil:
		return nil, err
	case !core.FieldAccepts(f.Codecs, r):
		return nil, fmt.Errorf("%s: codec not allowed here", tok)
	}
	return core.EncodeRecord(r)
}
