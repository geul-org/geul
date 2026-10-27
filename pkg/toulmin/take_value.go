//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeValue — read an embedded Value record at *pos and advance past it
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// TakeValue decodes the embedded record at ws[*pos], which must be a Value, and
// advances *pos by its Length.
func TakeValue(ws []core.Word, pos *int) (core.Record, error) {
	r, n, err := core.DecodeRecord(ws[*pos:])
	if err != nil {
		return core.Record{}, err
	}
	if r.Kind != core.KindValue {
		return core.Record{}, errors.New("geul/toulmin: embedded record is not a Value")
	}
	*pos += n
	return r, nil
}
