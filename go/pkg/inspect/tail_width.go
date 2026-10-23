//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what tailWidth — the words the fields after a spec's rest field take; false for several rest fields or a non-fixed field after one
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// tailWidth sums the widths of the fields after fs's rest field (edge/FIELD):
// at most one rest field, and every field after it a fixed raw field, since a
// record's end is found only by reading it from its start. Without a rest
// field it is 0.
func tailWidth(fs []codecspec.Field) (int, bool) {
	sum, rest := 0, false
	for _, f := range fs {
		switch {
		case f.Width == int(core.FieldRest) && rest:
			return 0, false
		case f.Width == int(core.FieldRest):
			rest = true
		case rest && f.Width == 0:
			return 0, false
		case rest:
			sum += f.Width
		}
	}
	return sum, true
}
