//ff:func feature=toulmin type=decoder control=sequence
//ff:what TakeVFix — read an embedded SIGNED-DECIMAL Value at *pos as a VFix
package toulmin

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// TakeVFix reads the embedded SIGNED-DECIMAL Value at ws[*pos] as a verdict or
// qualifier VFix and advances *pos.
func TakeVFix(ws []core.Word, pos *int) (*VFix, error) {
	r, err := TakeValue(ws, pos)
	if err != nil {
		return nil, err
	}
	v, err := DecodeVFixValue(r)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
