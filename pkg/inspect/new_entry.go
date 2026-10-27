//ff:func feature=inspect type=decoder control=sequence
//ff:what newEntry — a bare Entry: the record, its encoding and its core spec name (a slot marker by its marker name)
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// newEntry returns the uninterpreted Entry of r: the record, its words (a
// decoded record always re-encodes) and the core name of its codec, or of the
// slot marker it is (core.SlotMarkerName; misplaced as a record, which schema
// reports).
func newEntry(r core.Record) Entry {
	ws, _ := core.EncodeRecord(r)
	name := core.CodecName(r.Kind, r.Codec)
	if core.IsSlotMarker(r) {
		name = core.SlotMarkerName(r.Codec)
	}
	return Entry{Record: r, Words: ws, Name: name}
}
