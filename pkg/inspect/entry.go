//ff:type feature=inspect type=model
//ff:what Entry — one interpreted record: words, spec name, decoded content, Edge Id and referenced Edge Ids
package inspect

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// Entry is one interpreted record. Words is its encoding, TAG WORD included.
// Name is its spec name ("" when unknown). Data is the decoded content — Value,
// Widths, DocID, ImportDoc, Source or a toulmin TextPacket, ControlPacket,
// NodePacket or EdgePacket — and nil when the record has no lossless typed
// form. ID is the record's Edge Id and Refs the non-zero Edge Ids it
// references in field order; both are empty when Data carries none. Spec is
// the codec spec fields of the record's codec (edge/DECLARE), nil when undeclared.
type Entry struct {
	Record core.Record
	Words  []core.Word
	Name   string
	Data   any
	ID     uint64
	Refs   []uint64
	Spec   []codecspec.Field
}
