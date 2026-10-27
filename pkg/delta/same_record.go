//ff:func feature=delta type=helper control=sequence
//ff:what sameRecord — whether two records encode to the same words
package delta

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// sameRecord compares the encodings of a and b.
func sameRecord(a, b core.Record) bool {
	wa, _ := core.EncodeRecord(a)
	wb, _ := core.EncodeRecord(b)
	return slices.Equal(wa, wb)
}
