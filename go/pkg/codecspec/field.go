//ff:type feature=codecspec type=model
//ff:what Field — one declared payload field: name, width (0 = record), CODEC targets, enum names
package codecspec

import "github.com/geul-org/geul/go/pkg/core"

// Field is one declared payload field (edge/FIELD). Width 0 is a record field
// read by its own TAG Length; n is a raw field of n words. Codecs are its
// CODEC targets (core.TargetWord): the Kind·Codecs a record field may have
// (any when empty), or for a raw field the one codec its words are the
// Payload of. Enums maps a raw code in canonical hex (core.FormatID) to its
// BIND name.
type Field struct {
	Name   string
	Width  int
	Codecs []core.Word
	Enums  map[string]string
}
