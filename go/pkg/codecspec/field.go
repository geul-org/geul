//ff:type feature=codecspec type=model
//ff:what Field — one declared payload field: name, width (0 = record), expected TAG, enum names
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// Field is one declared payload field (edge/FIELD). Width 0 is a record field
// read by its own TAG Length; n is a raw field of n words. Expect is the
// Kind·Codec a record field must have when HasExpect. Enums maps a raw code in
// canonical hex (core.FormatID) to its BIND name.
type Field struct {
	Name      string
	Width     int
	Expect    core.Word
	HasExpect bool
	Enums     map[string]string
}
